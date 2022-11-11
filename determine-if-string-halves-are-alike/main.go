package main

import (
	"log"
	"strings"
)

func main() {
	var s string
	s = "book"
	log.Println(halvesAreAlike(s))

	s = "textbook"
	log.Println(halvesAreAlike(s))
}

func halvesAreAlike(s string) bool {
	a := strings.ToLower(s[:len(s)/2])
	b := strings.ToLower(s[len(s)/2:])

	return vowelCount(a) == vowelCount(b)
}

func vowelCount(a string) int {
	count := 0
	vowelsMap := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	for _, v := range a {
		if vowelsMap[v] {
			count++
		}
	}
	return count
}
