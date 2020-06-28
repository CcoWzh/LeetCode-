package main

import "fmt"

func uniquePaths(m int, n int) int {

	dp := make([][]int, m)
	for i := 0; i < n; i++ {
		dp[0] = append(dp[0], 1)
	}
	for i := 1; i < m; i++ {
		dp[i] = append(dp[i], 1)
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i] = append(dp[i], dp[i][j-1]+dp[i-1][j])
		}
	}

	fmt.Println(dp)
	return dp[m-1][n-1]
}

func main() {
	m := 2
	n := 1
	fmt.Println(uniquePaths(m, n))
}
