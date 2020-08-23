package main

import "fmt"

func minFlipsMonoIncr(S string) int {
	n := len(S)
	//dp[i][0]:以i结尾，最终状态为0的最小翻转次数
	//dp[i][1]:以i结尾，最终状态为1的最小翻转次数
	dp := make([][2]int, n)

	if S[0] == '0' {
		dp[0][0] = 0
		dp[0][1] = 1
	} else {
		dp[0][0] = 1
		dp[0][1] = 0
	}

	for i := 1; i < n; i++ {
		if S[i] == '0' {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + 1
		} else if S[i] == '1' {
			dp[i][0] = dp[i-1][0] + 1
			dp[i][1] = min(dp[i-1][0], dp[i-1][1])
		}
	}

	return min(dp[n-1][0], dp[n-1][1])
}

func min(a, b int) int {
	if a > b {
		a = b
	}
	return a
}

func main() {
	S := "00110"
	fmt.Println(minFlipsMonoIncr(S))
}
