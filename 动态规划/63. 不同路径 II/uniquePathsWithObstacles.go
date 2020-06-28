package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	//初始化dp
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dp[i] == nil {
				dp[i] = make([]int, n)
			}
			dp[i][j] = 1
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//遇到障碍物，标记为0，注意，我们是从i=0,j=0开始的
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			//第一行，如果遇到障碍物，则障碍物之后的路径都走不通
			if i == 0 && j > 0 {
				dp[i][j] = dp[i][j-1]
			}
			//第一列，如果遇到障碍物，则障碍物之后的路径都走不通
			if j == 0 && i > 0 {
				dp[i][j] = dp[i-1][j]
			}
			//如果遇到障碍物，来自障碍物的那个方向，没有路径输入，即0
			if i > 0 && j > 0 {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
		}
	}

	fmt.Println(dp)
	return dp[m-1][n-1]
}

func main() {
	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0}}
	fmt.Println(uniquePathsWithObstacles(obstacleGrid))
}
