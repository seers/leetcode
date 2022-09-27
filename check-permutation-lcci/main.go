package main

import (
	"fmt"
	"reflect"
)

func main() {
	s1 := "abc"
	s2 := "bca"
	fmt.Println(CheckPermutation(s1, s2))
	s1 = "abc"
	s2 = "bad"
	fmt.Println(CheckPermutation(s1, s2))
	s1 = "aa"
	s2 = "bb"
	fmt.Println(CheckPermutation(s1, s2))
	s1 = "abb"
	s2 = "aab"
	fmt.Println(CheckPermutation(s1, s2))
}

// 遍历字符串，构造rune和出现次数的键值对，比较两个map是否一致
func CheckPermutation(s1 string, s2 string) bool {
	s1map := make(map[rune]int)
	s2map := make(map[rune]int)
	for _, v := range s1 {
		s1map[v]++
	}
	for _, v := range s2 {
		s2map[v]++
	}
	return reflect.DeepEqual(s1map, s2map)
}
