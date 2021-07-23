package main

import "fmt"

func main() {

	s := []string{"one", "two", "three"}
	s1 := []string{"one", "two"}
	fmt.Println(max(s))
	fmt.Println(max(s1))
}

func max(s []string) string {

	var i, length, maxLen int
	longest := s[i]
	for _, word := range s {
		length = len(word)

		if length > maxLen {
			longest = word
		}
	}
	return longest
}
