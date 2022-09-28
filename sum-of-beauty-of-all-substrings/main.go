package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := "aabcb"
	fmt.Println(beautySum(s))

	s = "aabcbaa"
	fmt.Println(beautySum(s))
}

// 用切片保存，排序后最大值减去最小值
func beautySum(s string) int {
	substrs := getAllSubStr(s)
	sum := 0
	for _, substr := range substrs {
		arr := []int{}
		for _, strrune := range substr {
			arr = append(arr, strings.Count(substr, string(strrune)))
		}
		sort.Ints(arr)
		sum += arr[len(arr)-1] - arr[0]
	}
	return sum
}

// 获取所有子字符串
func getAllSubStr(s string) []string {

}
