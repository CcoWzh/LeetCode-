package main

import "fmt"

//300. 最长上升子序列
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	max_v := 1
	for i := 1; i < len(nums); i++ {
		temp := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				temp = max(temp, dp[j])
			}
		}
		dp[i] = temp + 1
		max_v = max(max_v, dp[i])
	}

	fmt.Println(dp)

	return max_v
}

func max(a, b int) int {
	if a < b {
		a = b
	}
	return a
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}
