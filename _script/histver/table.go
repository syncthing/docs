package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"regexp"
	"sort"
)

type tableRow struct {
	Version string
	Runtime string
	Date    string
}

func (r *tableRow) fromStrings(ss []string) error {
	if len(ss) < 3 {
		return fmt.Errorf("not enough fields")
	}
	r.Version = ss[0]
	r.Runtime = ss[1]
	r.Date = ss[2]
	return nil
}

func (r *tableRow) fromVersion(ver string) error {
	// syncthing v1.23.1-rc.1 "Fermium Flea" (go1.19.5 darwin-arm64) teamcity@build.syncthing.net 2023-01-12 03:30:17 UTC [stnoupgrade]
	exp := regexp.MustCompile(`syncthing (v\d+\.\d+\.\d+).*(go\d+\.\d+(?:\.\d+)?).*(\d{4}-\d{2}-\d{2}) `)
	m := exp.FindStringSubmatch(ver)
	if len(m) < 3 {
		return fmt.Errorf("failed to parse version")
	}
	r.Version = m[1]
	r.Runtime = m[2]
	r.Date = m[3]
	return nil
}

func (r *tableRow) toStrings() []string {
	return []string{r.Version, r.Runtime, r.Date}
}

var tableHeader = []string{"Version", "Runtime", "Date"}

func writeTable(w io.Writer, rows []*tableRow) error {
	sort.Slice(rows, func(a, b int) bool {
		if rows[a].Date == rows[b].Date {
			return rows[a].Version > rows[b].Version
		}
		return rows[a].Date > rows[b].Date
	})
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

func readTable(r io.Reader) ([]*tableRow, error) {
	cr := csv.NewReader(r)
	var rows []*tableRow
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
		rows = append(rows, &row)
	}
	return rows, nil
}
