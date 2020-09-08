package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])
	dp := make([][]int, n)
	maxSide := 0
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		//其实，只需要对于第一行和第一列的数，原封不动
		for j := 0; j < m; j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if matrix[i][j] == '1' {
				maxSide = 1
			}
		}
	}
	//从坐标(1,1)开始遍历
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == '1' {
				//左上，左边和上边
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			} else {
				dp[i][j] = 0
			}
			//更新结果
			if maxSide < dp[i][j] {
				maxSide = dp[i][j]
			}
		}
	}

	return maxSide * maxSide
}

func min(a, b int) int {
	if a > b {
		a = b
	}
	return a
}

func main() {
	matrix := [][]byte{
		{'1', '1'},
		{'1', '1'}}
	fmt.Println(maximalSquare(matrix))
}
