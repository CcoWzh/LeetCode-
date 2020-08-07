package main

import (
	"fmt"
)

//func maxProfit(prices []int) int {
//	if len(prices) == 0 {
//		return 0
//	}
//	profit := make([]int, len(prices))
//
//	for i := 0; i < len(prices); i++ {
//		min := prices[i]
//		max := prices[i]
//		for j := i; j < len(prices); j++ {
//			if prices[j] > max {
//				max = prices[j]
//			}
//		}
//		profit[i] = max - min
//	}
//	sort.Ints(profit)
//	fmt.Println(profit)
//
//	return profit[len(profit)-1]
//}

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	//nums := []int{7, 6, 4, 3, 1}
	//nums := []int{}
	fmt.Println(maxProfit(nums))
}

func maxProfit(prices []int) int {
	minPrice := int(1e9)
	maxPrice := 0
	for _, v := range prices {
		maxPrice = max(v-minPrice, maxPrice)
		minPrice = min(v, minPrice)
	}
	return maxPrice
}

func max(a, b int) int {
	if a < b {
		a = b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		a = b
	}
	return a
}
