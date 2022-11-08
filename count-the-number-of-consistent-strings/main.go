package main

import (
	"fmt"
	"strings"
)

func main() {
	var allowed string
	var words []string

	allowed = "ab"
	words = []string{"ad", "bd", "aaab", "baa", "badab"}
	fmt.Println(countConsistentStrings(allowed, words))

	allowed = "abc"
	words = []string{"a", "b", "c", "ab", "ac", "bc", "abc"}
	fmt.Println(countConsistentStrings(allowed, words))

	allowed = "cad"
	words = []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}
	fmt.Println(countConsistentStrings(allowed, words))
}

func countConsistentStrings(allowed string, words []string) int {
	count := 0
	for _, word := range words {
		for key, char := range word {
			if !strings.ContainsRune(allowed, char) {
				break
			}
			if key == len(word)-1 {
				count++
			}
		}
	}
	return count
}
