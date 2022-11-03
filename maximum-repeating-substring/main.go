package main

import (
	"fmt"
	"strings"
)

func main() {
	sequence := "ababc"
	word := "ab"
	fmt.Println(maxRepeating(sequence, word))

	sequence = "ababc"
	word = "ba"
	fmt.Println(maxRepeating(sequence, word))

	sequence = "ababc"
	word = "ab"
	fmt.Println(maxRepeating(sequence, word))
}

func maxRepeating(sequence string, word string) int {
	k := 0
	for {
		if strings.Count(sequence, strings.Repeat(word, k)) == 0 {
			break
		}
		k++
	}
	return k - 1
}
