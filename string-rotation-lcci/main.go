package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "aba"
	s2 := ""
	fmt.Println(isFlipedString(s1, s2))
}

func isFlipedString(s1 string, s2 string) bool {
	if s1 == "" && s2 == "" {
		return true
	}
	if s1 != "" && s2 == "" {
		return false
	}
	return strings.Contains(strings.Repeat(s1, 2), s2)
}
