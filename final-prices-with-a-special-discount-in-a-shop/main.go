package main

import "fmt"

func main() {
	prices := []int{8, 4, 6, 2, 3}
	fmt.Println(finalPrices(prices))

	prices = []int{1, 2, 3, 4, 5}
	fmt.Println(finalPrices(prices))

	prices = []int{10, 1, 1, 6}
	fmt.Println(finalPrices(prices))
}

// 遍历数组，如果小就减去
func finalPrices(prices []int) []int {
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				prices[i] -= prices[j]
				break
			}
		}
	}
	return prices
}
