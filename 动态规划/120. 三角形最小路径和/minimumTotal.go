package main

import "fmt"

//120. 三角形最小路径和
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}
	//dp[i][j]表示从(0,0)--->(i,j)的最小路径和
	dp := make([][]int, n)
	//初始化数组
	for i := 0; i < n; i++ {
		m := len(triangle[i])
		for j := 0; j < m; j++ {
			dp[i] = append(dp[i], 0)
		}
	}
	dp[0][0] = triangle[0][0]
	//动态规划
	for i := 1; i < n; i++ {
		m := len(triangle[i])
		for j := 0; j < m; j++ {
			if j-1 < 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j >= len(triangle[i-1]) {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			}
		}
	}
	//求最小值
	min := dp[n-1][0]
	for i := 0; i < len(triangle[n-1]); i++ {
		if min > dp[n-1][i] {
			min = dp[n-1][i]
		}
	}

	return min
}

func min(a, b int) int {
	if a > b {
		a = b
	}
	return a
}

func main() {
	//triangle := [][]int{
	//	{2},
	//	{3, 4},
	//	{6, 5, 7},
	//	{4, 1, 8, 3},
	//}
	triangle := [][]int{}
	fmt.Println(minimumTotal(triangle))
}
