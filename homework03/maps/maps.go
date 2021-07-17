package main

import (
	"fmt"
	"sort"
)

func main() {
	m1 := map[int]string{2: "a", 0: "b", 1: "c"}
	m2 := map[int]string{10: "aa", 0: "bb", 500: "cc"}

	printSorted(m1)
	printSorted(m2)
}

func printSorted(m map[int]string) {

	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		fmt.Println(m[k])
	}
}
