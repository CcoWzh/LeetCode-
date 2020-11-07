package main

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)
	if n == 0 || m == 0 {
		return 0
	}
	//dp[i][j]表示text1[0~i]和text2[0~j]的最长公共子序列
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ { //两种状态，等于或者不等于
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[n][m]
}

func max(a, b int) int {
	if a < b {
		a = b
	}
	return a
}

func main() {
	text1 := "abc"
	text2 := "abe"
	fmt.Println(longestCommonSubsequence(text1, text2))
}
