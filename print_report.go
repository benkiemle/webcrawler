package main

import (
	"fmt"
	"sort"
)

type kv struct {
	Key   string
	Value int
}

func (cfg *config) printReport() {
	fmt.Println("=============================")
	fmt.Printf("  REPORT for %s\n", cfg.baseURL.String())
	fmt.Println("=============================")

	sorted := sortPages(cfg.pages)
	for _, s := range sorted {
		fmt.Printf("Found %d internal links to %s\n", s.Value, s.Key)
	}
}

func sortPages(pages map[string]int) []kv {

	var ss []kv
	for k, v := range pages {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		if ss[i].Value == ss[j].Value {
			return ss[i].Key < ss[j].Key
		}
		return ss[i].Value > ss[j].Value
	})

	return ss
}
