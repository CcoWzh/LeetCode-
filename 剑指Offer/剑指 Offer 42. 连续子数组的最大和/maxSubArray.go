package main

import "fmt"

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	//dp[i] 表示以i结尾的数，
	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
	}

	max := dp[0]
	for i := 0; i < n; i++ {
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
