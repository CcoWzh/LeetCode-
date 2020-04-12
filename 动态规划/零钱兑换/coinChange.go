package main

import "fmt"

//322. 零钱兑换
//F(i) 为组成金额 i 所需最少的硬币数量
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount+1
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}

func min(a, b int) int {
	if a > b {
		a = b
	}
	return a
}

func main() {
	coins := []int{1, 2, 5}
	amount := 0
	fmt.Println(coinChange(coins, amount))
}
