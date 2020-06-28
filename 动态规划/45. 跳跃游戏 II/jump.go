package main

import "fmt"

//45. 跳跃游戏 II
func jump(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = n + 1
	}

	dp[n-1] = 0
	//动态规划，从前往后
	//dp[i]表示从i位置跳到最后的位置，需要的最小步数
	for i := n - 1; i >= 0; i-- {
		//dp[i]=min(dp[i+1]......dp[i+nums[i])+1
		for j := i + 1; j <= i+nums[i]; j++ {
			if j >= n {
				break
			}
			dp[i] = min(dp[j]+1, dp[i])
		}
	}
	//fmt.Println(dp)

	return dp[0]
}

func min(a, b int) int {
	if a > b {
		a = b
	}
	return a
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums))
}
