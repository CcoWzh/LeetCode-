package main

import (
	"fmt"
)

//动态规划，找零钱
func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 0; i-j*j >= 0; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	fmt.Println(dp)
	return dp[n]
}

func min(a, b int) int {
	if a > b {
		a = b
	}
	return a
}

func main() {
	n := 12
	fmt.Println(numSquares(n))
}
