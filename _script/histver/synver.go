package main

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

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
	var table []*tableRow
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

	seen := make(map[string]struct{})
	for _, row := range table {
		seen[row.Version] = struct{}{}
	}

	// Get version information for all releases not yet in the versions
	// table.
	for _, rel := range releases {
		if _, ok := seen[*rel.TagName]; ok {
			continue
		}
		log.Println("Checking", *rel.TagName)
		if row, err := getReleaseVersion(rel); err != nil {
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

func getReleaseVersion(rel *github.RepositoryRelease) (*tableRow, error) {
	goos := runtime.GOOS
	if goos == "darwin" {
		goos = "macos"
	}
	find := fmt.Sprintf("syncthing-%s-%s", goos, runtime.GOARCH)
	for _, asset := range rel.Assets {
		if strings.HasPrefix(*asset.Name, find) {
			log.Println("Downloading", *asset.Name)
			resp, err := http.Get(*asset.BrowserDownloadURL)
			if err != nil {
				return nil, err
			}
			bs, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				return nil, err
			}
			switch filepath.Ext(*asset.Name) {
			case ".zip":
				return getReleaseVersionZip(bs)
			default:
				return getReleaseVersionTarGz(bs)
			}
		}
	}
	return nil, fmt.Errorf("no asset found")
}

func getReleaseVersionZip(bs []byte) (*tableRow, error) {
	zr, err := zip.NewReader(bytes.NewReader(bs), int64(len(bs)))
	if err != nil {
		return nil, err
	}
	for _, f := range zr.File {
		if strings.Contains(path.Dir(f.Name), "/") {
			// Skip files not at top level
			continue
		}
		if path.Base(f.Name) != "syncthing" {
			continue
		}
		rd, err := f.Open()
		if err != nil {
			return nil, err
		}

		return getVersionFromReader(rd)
	}
	return nil, fmt.Errorf("no syncthing binary found")
}

func getReleaseVersionTarGz(bs []byte) (*tableRow, error) {
	gr, err := gzip.NewReader(bytes.NewReader(bs))
	if err != nil {
		return nil, err
	}
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			break
		}
		if path.Base(hdr.Name) != "syncthing" {
			continue
		}

		return getVersionFromReader(tr)
	}
	return nil, fmt.Errorf("no syncthing binary found")
}

func getVersionFromReader(r io.Reader) (*tableRow, error) {
	fd, err := os.CreateTemp("", "syncthing")
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(fd, r); err != nil {
		return nil, err
	}
	fd.Close()
	defer os.Remove(fd.Name())
	if err := os.Chmod(fd.Name(), 0o755); err != nil {
		return nil, err
	}

	return getVersionFromCommand(fd.Name())
}

func getVersionFromCommand(name string) (*tableRow, error) {
	cmd := exec.Command(name, "--version")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var r tableRow
	if err := r.fromVersion(string(out)); err != nil {
		return nil, err
	}
	return &r, nil
}
