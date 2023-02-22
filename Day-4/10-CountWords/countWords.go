package main

// Follow TDD to count the number of words in a given sentence.

import "strings"

func CountWords(s string) int {
	if s != "" {
		return len(strings.Split(s, " "))
	}
	return 0
}
