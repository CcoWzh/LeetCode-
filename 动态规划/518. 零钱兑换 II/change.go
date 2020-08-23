package main

import "fmt"

func change(amount int, coins []int) int {
	n := len(coins)
	//dp[i][j]: 前i个物品凑成amount=j时的方法总数
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1 //当amount=0时，只有一种方法，即，什么都不选
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] >= 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]] //不选+选(选的时候，就不是i-1了)
			} else {
				dp[i][j] = dp[i-1][j] //不选
			}
		}
	}

	fmt.Println(dp)
	return dp[n][amount]
}

func main() {
	amount := 5
	coins := []int{1,2,5}
	fmt.Println(change(amount, coins))
}
