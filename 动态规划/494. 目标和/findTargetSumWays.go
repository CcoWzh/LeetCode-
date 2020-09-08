package main

import "fmt"

func findTargetSumWays(nums []int, S int) int {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	if sum < S { //非负整数数组,不需要绝对值
		//目标值比数组和还要大，这是绝对凑不出来的
		return 0
	}
	target := sum + S
	if target%2 != 0 {
		return 0
	}
	target = target / 2
	//0-1背包问题，从数组nums中，选出若干个数，使得这些数加起来恰好等于target
	//求这样的组合有多少种
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, target+1)
	}
	dp[0][0] = 1 //需要注意的是，初始值的问题
	//dp[i][j]表示当数组只有i个数字时，可以凑成和为j的子数组的个数
	for i := 1; i <= n; i++ {
		for j := 0; j <= target; j++ {
			if j >= nums[i-1] {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
			} else {
				dp[i][j] = dp[i-1][j] //只能不选了
			}
		}
	}
	//fmt.Println(dp)
	return dp[n][target]
}

func main() {
	nums := []int{1, 2, 3, 4}
	S := 6
	fmt.Println(findTargetSumWays(nums, S))
}
