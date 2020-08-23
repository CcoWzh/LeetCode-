package main

import "fmt"

/**
给一个可装载重量为sum/2的背包和N个物品，每个物品的重量为nums[i]。现在让你装物品，是否存在一种装法，能够恰好将背包装满？
*/
func canPartition(nums []int) bool {
	//0-1背包问题
	n, sum := len(nums), 0
	if n == 0 {
		return false
	}
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum%2 == 1 { //和为奇数时，不可能划分成两个和相等的集合
		return false
	}
	target := sum / 2
	//题目转换成，能不能从nums中，选出若干个数，使得这些数加起来等于target
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, target+1)
		dp[i][0] = true
	}
	//比如说，如果dp[4][9] = true，其含义为：对于容量为 9 的背包，若只是用前 4 个物品，可以有一种方法把背包恰好装满。
	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			if j-nums[i-1] < 0 { //背包容量为j，装不下当前值时
				dp[i][j] = dp[i-1][j]
			} else { //背包容量为j，装得下当前值时，分为两种情况：不装 || 装
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}

	return dp[n][target]
}

func main() {
	nums := []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
}
