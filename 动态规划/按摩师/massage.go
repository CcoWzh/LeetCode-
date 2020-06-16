package main

import (
	"fmt"
)

//面试题 17.16. 按摩师
func massage(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//dp0,dp1分别表示dp[i-1][0],de[i-1][1]
	dp0, dp1 := 0, nums[0]
	//从1开始
	for i := 1; i < len(nums); i++ {
		temp0 := max(dp0, dp1)
		temp1 := dp0 + nums[i]

		dp0 = temp0
		dp1 = temp1
	}
	return max(dp1, dp0)
}

func max(a, b int) int {
	if a < b {
		a = b
	}
	return a
}

func main() {
	nums := []int{1, 1}
	fmt.Println(massage(nums))
}
