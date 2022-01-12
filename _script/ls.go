package main

import (
	"encoding/json"
	"log"
	"os"
	"sort"
)

func main() {
	dir := os.Args[1]
	if err := ls(dir); err != nil {
		log.Fatalln("ls:", err)
	}
}

func ls(dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		names = append(names, entry.Name())
	}
	sort.Strings(names)

	enc := json.NewEncoder(os.Stdout)
	return enc.Encode(map[string][]string{"entries": names})
}
