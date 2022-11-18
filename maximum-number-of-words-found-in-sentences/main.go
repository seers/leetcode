package main

import (
	"fmt"
	"strings"
)

func main() {
	var sentences []string
	sentences = []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	fmt.Println(mostWordsFound(sentences))

	sentences = []string{"please wait", "continue to fight", "continue to win"}
	fmt.Println(mostWordsFound(sentences))
}

func mostWordsFound(sentences []string) int {
	max := 0
	for _, sentence := range sentences {
		num := len(strings.Split(sentence, " "))
		if num > max {
			max = num
		}
	}
	return max
}
