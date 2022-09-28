package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}

// 遍历的同时将之前的index记录下来，然后遇到target-v满足时候返回
func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for k, v := range nums {
		if p, ok := hashTable[target-v]; ok {
			return []int{p, k}
		}
		hashTable[v] = k
	}
	return nil
}
