package main

import (
	"fmt"
	"strings"
)

func main() {
	word1 := []string{"ab", "c"}
	word2 := []string{"a", "bc"}
	fmt.Println(arrayStringsAreEqual(word1, word2))

	word1 = []string{"a", "cb"}
	word2 = []string{"ab", "c"}
	fmt.Println(arrayStringsAreEqual(word1, word2))

	word1 = []string{"abc", "d", "defg"}
	word2 = []string{"abcddefg"}
	fmt.Println(arrayStringsAreEqual(word1, word2))
}

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	return strings.Compare(strings.Join(word1, ""), strings.Join(word2, "")) == 0
}
