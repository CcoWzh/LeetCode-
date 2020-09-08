package main

import "fmt"

func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		curPrice := prices[i] - prices[i-1]
		if curPrice > 0 {
			res += curPrice
		}
	}

	return res
}

func main() {
	prices := []int{1, 9, 6, 9, 1, 7, 1, 1, 5, 9, 9, 9}
	fmt.Println(maxProfit(prices))
}
