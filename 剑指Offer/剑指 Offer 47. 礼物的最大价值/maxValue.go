package main

import "fmt"

/**动态规划
状态转移方程:dp[i][j] = grid[i][j] + max(dp[i-1][j], dp[i][j-1])
 */
func maxValue(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}

	dp[0][0] = grid[0][0]
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j > 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			}
			if j == 0 && i > 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			}
			if i > 0 && j > 0 {
				dp[i][j] = grid[i][j] + max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[n-1][m-1]
}

func max(a, b int) int {
	if a < b {
		a = b
	}
	return a
}

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1}}
	fmt.Println(maxValue(grid))
}
