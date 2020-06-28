package main

import "fmt"

//300.最长上升子序列
/**
dp[i]表示以最大为nums[i]结尾的最长上升子序列的数量
我们只要找到前面那些结尾比 nums[i] 小的子序列，然后把 nums[i] 接到最后，
就可以形成一个新的递增子序列，而且这个新的子序列长度+1
这个例子告诉我，动态规划不一定
 */
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
