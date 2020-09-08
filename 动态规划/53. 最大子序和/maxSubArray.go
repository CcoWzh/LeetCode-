package main

import "fmt"

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -2147483648
	}
	dp := make([]int, n)
	//dp[i] 表示以i结尾的，连续子数组的最大和
	//dp[i] = max(dp[i-1]+nums[i], nums[i])
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
	}
	// fmt.Println(dp)
	res := dp[0]
	for i := 0; i < n; i++ {
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}

func max(a, b int) int {
	if a < b {
		a = b
	}
	return a
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
