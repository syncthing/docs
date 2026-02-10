package main

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"cmp"
	"compress/gzip"
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/google/go-github/v49/github"
)

func main() {
	versionsFile := flag.String("file", "versions.csv", "Path to versions CSV file")
	flag.Parse()

	// Load all known releases
	ctx := context.Background()
	releases, err := getReleases(ctx)
	if err != nil {
		log.Fatalln("Listing GitHub releases:", err)
	}

	// Load current versions table
	var table []tableRow
	fd, err := os.Open(*versionsFile)
	if os.IsNotExist(err) {
		// File doesn't exist yet. That's alright.
	} else if err != nil {
		log.Fatalln("Reading existing versions:", err)
	} else {
		table, err = readTable(fd)
		if err != nil {
			log.Fatalln("Reading existing versions:", err)
		}
	}

	seen := make(map[string]tableRow)
	for _, row := range table {
		seen[row.Version] = row
	}

	// Get version information for all releases not yet in the versions
	// table.
	for _, rel := range releases {
		row := seen[rel.GetTagName()]
		if row.complete() {
			continue
		}
		log.Println("Checking", rel.GetTagName())
		if row, err := getReleaseVersion(rel, row); err != nil {
			log.Printf("%s: %v", *rel.TagName, err)
		} else {
			table = append(table, row)
		}
	}

	// Save a new versions table.
	tw, err := os.Create(*versionsFile)
	if err != nil {
		log.Fatalln("Creating versions table:", err)
	}
	if err := writeTable(tw, table); err != nil {
		log.Fatalln("Writing versions table:", err)
	}
	if err := tw.Close(); err != nil {
		log.Fatalln("Writing versions table:", err)
	}
}

func getReleases(ctx context.Context) ([]*github.RepositoryRelease, error) {
	client := github.NewClient(nil)
	opts := &github.ListOptions{
		PerPage: 100,
	}

	var releases []*github.RepositoryRelease
	for {
		rels, resp, err := client.Repositories.ListReleases(ctx, "syncthing", "syncthing", opts)
		if err != nil {
			return nil, err
		}
		for _, rel := range rels {
			if *rel.Prerelease {
				continue
			}
			releases = append(releases, rel)
		}
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	sort.Slice(releases, func(a, b int) bool {
		return releases[a].GetPublishedAt().After(releases[b].GetPublishedAt().Time)
	})
	return releases, nil
}

func getReleaseVersion(rel *github.RepositoryRelease, row tableRow) (tableRow, error) {
	row = tableRow{
		Version:        cmp.Or(row.Version, rel.GetTagName()),
		LinuxRuntime:   row.LinuxRuntime,
		WindowsRuntime: row.WindowsRuntime,
		MacOSRuntime:   row.MacOSRuntime,
		Date:           cmp.Or(row.Date, rel.GetPublishedAt().Format("2006-01-01")),
	}

	var osarchs []string
	if row.LinuxRuntime == "" {
		osarchs = append(osarchs, "linux-amd64")
	}
	if row.WindowsRuntime == "" {
		osarchs = append(osarchs, "windows-amd64")
	}
	if row.MacOSRuntime == "" {
		osarchs = append(osarchs, "macos-arm64", "macos-amd64", "darwin-amd64")
	}

	if len(osarchs) == 0 {
		return row, nil
	}

	var mut sync.Mutex
	var wg sync.WaitGroup
	for _, osarch := range osarchs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tr, err := getReleaseRow(rel, osarch)
			mut.Lock()
			if err != nil {
				log.Printf("%s %s: %s", rel.GetTagName(), osarch, err)
			} else {
				row = row.merge(tr)
			}
			mut.Unlock()
		}()
	}
	wg.Wait()

	return row, nil
}

func getReleaseRow(rel *github.RepositoryRelease, osarch string) (tableRow, error) {
	find := "syncthing-" + osarch
	for _, asset := range rel.Assets {
		if strings.HasPrefix(*asset.Name, find) {
			log.Println("Downloading", *asset.Name)
			resp, err := http.Get(*asset.BrowserDownloadURL)
			if err != nil {
				return tableRow{}, err
			}
			bs, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				return tableRow{}, err
			}
			switch filepath.Ext(*asset.Name) {
			case ".zip":
				return getReleaseVersionZip(bs, osarch)
			default:
				return getReleaseVersionTarGz(bs, osarch)
			}
		}
	}
	return tableRow{}, errors.New("no asset found")
}

func getReleaseVersionZip(bs []byte, osarch string) (tableRow, error) {
	zr, err := zip.NewReader(bytes.NewReader(bs), int64(len(bs)))
	if err != nil {
		return tableRow{}, err
	}
	for _, f := range zr.File {
		if strings.Contains(path.Dir(f.Name), "/") {
			// Skip files not at top level
			continue
		}
		if path.Base(f.Name) != "syncthing" && path.Base(f.Name) != "syncthing.exe" {
			continue
		}
		rd, err := f.Open()
		if err != nil {
			return tableRow{}, err
		}

		return getVersionFromReader(rd, osarch)
	}
	return tableRow{}, fmt.Errorf("no syncthing binary found")
}

func getReleaseVersionTarGz(bs []byte, osarch string) (tableRow, error) {
	gr, err := gzip.NewReader(bytes.NewReader(bs))
	if err != nil {
		return tableRow{}, err
	}
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			break
		}
		if strings.Contains(path.Dir(hdr.Name), "/") {
			// Skip files not at top level
			continue
		}
		if path.Base(hdr.Name) != "syncthing" {
			continue
		}

		return getVersionFromReader(tr, osarch)
	}
	return tableRow{}, fmt.Errorf("no syncthing binary found")
}

func getVersionFromReader(r io.Reader, osarch string) (tableRow, error) {
	fd, err := os.CreateTemp("", "syncthing")
	if err != nil {
		return tableRow{}, err
	}
	if _, err := io.Copy(fd, r); err != nil {
		return tableRow{}, err
	}
	fd.Close()
	defer os.Remove(fd.Name())
	if err := os.Chmod(fd.Name(), 0o755); err != nil {
		return tableRow{}, err
	}

	var errs []error
	if row, err := getVersionFromCommand(fd.Name(), osarch); err == nil {
		return row, nil
	} else {
		errs = append(errs, err)
	}
	if row, err := getVersionFromGo(fd.Name(), osarch); err == nil {
		return row, nil
	} else {
		errs = append(errs, err)
	}
	return tableRow{}, errors.Join(errs...)
}

func getVersionFromGo(name string, osarch string) (tableRow, error) {
	cmd := exec.Command("go", "version", "-m", name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return tableRow{}, fmt.Errorf("go version -m: %w: %s", err, out)
	}

	// % go version -m ~/bin/syncthing
	// /Users/jb/bin/syncthing: go1.25.7
	// path	github.com/syncthing/syncthing/cmd/syncthing
	// ...

	if idx := bytes.Index(out, []byte{'\n'}); idx < 0 {
		return tableRow{}, errors.New("go version -m: no version")
	} else {
		out = out[:idx]
	}

	if idx := bytes.LastIndex(out, []byte{' '}); idx < 0 {
		return tableRow{}, errors.New("go version -m: no version")
	} else {
		return runtimeRow(string(out[idx+1:]), osarch), nil
	}
}

func getVersionFromCommand(name, osarch string) (tableRow, error) {
	cmd := exec.Command(name, "--version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		cmd = exec.Command(name, "version")
		out, err = cmd.CombinedOutput()
	}
	if err != nil {
		return tableRow{}, fmt.Errorf("syncthing: running: %w: %s", err, out)
	}

	var r tableRow
	if err := r.fromVersion(string(out), osarch); err != nil {
		return tableRow{}, fmt.Errorf("syncthing: parse version: %w: %s", err, out)
	}
	return r, nil
}

func runtimeRow(version, osarch string) tableRow {
	switch {
	case strings.HasPrefix(osarch, "linux"):
		return tableRow{LinuxRuntime: version}
	case strings.HasPrefix(osarch, "windows"):
		return tableRow{WindowsRuntime: version}
	case strings.HasPrefix(osarch, "macos"), strings.HasPrefix(osarch, "darwin"):
		return tableRow{MacOSRuntime: version}
	default:
		return tableRow{}
	}
}
