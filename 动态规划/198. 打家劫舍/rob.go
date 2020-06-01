package main

import "fmt"

/*打家劫舍
dp[i]表示当前以i结尾的可获得的最大金额
状态转移方程：
dp[i] = max(dp[i-1], dp[i-2]+nums[i])
 */
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	if n == 2 {
		return max(nums[0], nums[1])
	} else if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	//max(nums[0], nums[1])这一步没有想到
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	fmt.Println(dp, dp[len(dp)-1])
	return dp[len(dp)-1]
}

func max(a, b int) int {
	if a < b {
		a = b
	}
	return a
}

func main() {
	nums := []int{2, 1, 1, 2}
	rob(nums)
}
