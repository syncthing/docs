package main

import (
	"encoding/json"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	dir := os.Args[1]
	if err := lsver(dir); err != nil {
		log.Fatalln("ls:", err)
	}
}

func lsver(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	versionExp := regexp.MustCompile(`^v[0-9]+\.[0-9]+\.[0-9]+$`)

	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		if versionExp.MatchString(entry.Name()) {
			names = append(names, entry.Name())
		}
	}
	sort.Slice(names, func(a, b int) bool {
		return compareVersions(names[a], names[b]) < 0
	})

	enc := json.NewEncoder(os.Stdout)
	return enc.Encode(map[string][]string{"entries": names})
}

func compareVersions(a, b string) int {
	a = strings.TrimPrefix(a, "v")
	b = strings.TrimPrefix(b, "v")
	as := strings.Split(a, ".")
	bs := strings.Split(b, ".")
	an := make([]int, len(as))
	bn := make([]int, len(bs))
	for i, v := range as {
		an[i], _ = strconv.Atoi(v)
	}
	for i, v := range bs {
		bn[i], _ = strconv.Atoi(v)
	}
	for i := 0; i < len(an) && i < len(bn); i++ {
		switch {
		case an[i] < bn[i]:
			return -1
		case an[i] > bn[i]:
			return 1
		}
	}
	switch {
	case len(an) < len(bn):
		return -1
	case len(an) > len(bn):
		return 1
	}
	return 0
}
