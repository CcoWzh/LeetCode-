package main

import "fmt"

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][]int, 3)
	for i := 0; i < 3; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = -prices[0]
	// dp[0][i]: 手上持有股票的最大收益
	// dp[1][i]: 手上不持有股票，并且不在冷冻期中的累计最大收益
	// dp[2][i]: 手上不持有股票，并且处于冷冻期中的累计最大收益
	for i := 1; i < n; i++ {
		dp[0][i] = max(dp[0][i-1], dp[1][i-1]-prices[i])
		dp[1][i] = max(dp[1][i-1], dp[2][i-1])
		dp[2][i] = dp[0][i-1] + prices[i]
	}
	fmt.Println(dp)

	return max(dp[1][n-1], dp[2][n-1])
}

func max(a, b int) int {
	if a < b {
		a = b
	}
	return a
}

func main() {
	prices := []int{1, 2, 3, 4, 3, 0, 5}
	fmt.Println(maxProfit(prices))
}
