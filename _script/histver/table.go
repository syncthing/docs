package main

import (
	"cmp"
	"encoding/csv"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"
)

type tableRow struct {
	Version        string
	LinuxRuntime   string
	WindowsRuntime string
	MacOSRuntime   string
	Date           string
}

func (t tableRow) merge(other tableRow) tableRow {
	return tableRow{
		Version:        cmp.Or(other.Version, t.Version),
		LinuxRuntime:   cmp.Or(other.LinuxRuntime, t.LinuxRuntime),
		WindowsRuntime: cmp.Or(other.WindowsRuntime, t.WindowsRuntime),
		MacOSRuntime:   cmp.Or(other.MacOSRuntime, t.MacOSRuntime),
		Date:           cmp.Or(other.Date, t.Date),
	}
}

func (t tableRow) complete() bool {
	return t.Version != "" && t.Date != "" && t.LinuxRuntime != "" && t.WindowsRuntime != "" && t.MacOSRuntime != ""
}

func (r *tableRow) fromStrings(ss []string) error {
	switch len(ss) {
	case 3:
		r.Version = strings.Trim(ss[0], "*")
		r.LinuxRuntime = strings.Trim(ss[1], "*")
		r.Date = strings.Trim(ss[2], "*")
	case 5:
		r.Version = strings.Trim(ss[0], "*")
		r.LinuxRuntime = strings.Trim(ss[1], "*")
		r.WindowsRuntime = strings.Trim(ss[2], "*")
		r.MacOSRuntime = strings.Trim(ss[3], "*")
		r.Date = strings.Trim(ss[4], "*")
	default:
		return fmt.Errorf("bad number of fields %d", len(ss))
	}
	return nil
}

func (r *tableRow) fromVersion(ver, osarch string) error {
	// syncthing v1.23.1-rc.1 "Fermium Flea" (go1.19.5 darwin-arm64) teamcity@build.syncthing.net 2023-01-12 03:30:17 UTC [stnoupgrade]
	exp := regexp.MustCompile(`syncthing (v\d+\.\d+\.\d+).*(go\d+\.\d+(?:\.\d+)?).*(\d{4}-\d{2}-\d{2}) `)
	m := exp.FindStringSubmatch(ver)
	if len(m) < 3 {
		return fmt.Errorf("failed to parse version")
	}
	r.Version = m[1]
	switch {
	case strings.HasPrefix(osarch, "linux"):
		r.LinuxRuntime = m[2]
	case strings.HasPrefix(osarch, "windows"):
		r.WindowsRuntime = m[2]
	case strings.HasPrefix(osarch, "macos"), strings.HasPrefix(osarch, "darwin"):
		r.MacOSRuntime = m[2]
	}
	r.Date = m[3]
	return nil
}

func (r tableRow) toStrings() []string {
	return []string{r.Version, r.LinuxRuntime, r.WindowsRuntime, r.MacOSRuntime, r.Date}
}

var tableHeader = []string{"Version", "Linux Runtime", "Windows Runtime", "macOS Runtime", "Date"}

func writeTable(w io.Writer, rows []tableRow) error {
	sort.Slice(rows, func(a, b int) bool {
		if rows[a].Date == rows[b].Date {
			return rows[a].Version > rows[b].Version
		}
		return rows[a].Date > rows[b].Date
	})

	prevSynMinor := ""
	for i := len(rows) - 1; i >= 0; i-- {
		r := &rows[i]
		// Bold major/minor Syncthing releases
		synMinor := r.Version[:strings.LastIndex(r.Version, ".")]
		if synMinor != prevSynMinor {
			prevSynMinor = synMinor
			r.Version = fmt.Sprintf("**%s**", r.Version)
		}
		// Bold runtime differences
		if r.LinuxRuntime != r.WindowsRuntime || r.WindowsRuntime != r.MacOSRuntime {
			r.LinuxRuntime = fmt.Sprintf("**%s**", r.LinuxRuntime)
			r.WindowsRuntime = fmt.Sprintf("**%s**", r.WindowsRuntime)
			r.MacOSRuntime = fmt.Sprintf("**%s**", r.MacOSRuntime)
		}
	}
	cw := csv.NewWriter(w)
	if err := cw.Write(tableHeader); err != nil {
		return err
	}
	for _, r := range rows {
		if err := cw.Write(r.toStrings()); err != nil {
			return err
		}
	}
	cw.Flush()
	return cw.Error()
}

func readTable(r io.Reader) ([]tableRow, error) {
	cr := csv.NewReader(r)
	var rows []tableRow
	for {
		ss, err := cr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if len(ss) == 0 {
			continue
		}
		if ss[0] == tableHeader[0] {
			continue
		}
		var row tableRow
		if err := row.fromStrings(ss); err != nil {
			return nil, err
		}
		rows = append(rows, row)
	}
	return rows, nil
}
