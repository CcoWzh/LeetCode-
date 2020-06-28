package main

import (
	"fmt"
	"math"
)

/**
给定一个包含非负整数的 m x n 网格，
请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
 */
func minPathSum(grid [][]int) int {
	//if len(grid) == 0 {
	//	return 0
	//}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				grid[i][j] = grid[i][j]
			}
			if i == 0 && j > 0 {
				grid[i][j] = grid[i][j] + grid[i][j-1]
			}
			if i > 0 && j == 0 {
				grid[i][j] = grid[i][j] + grid[i-1][j]
			}
			if i > 0 && j > 0 {
				grid[i][j] = grid[i][j] + int(math.Min(float64(grid[i-1][j]), float64(grid[i][j-1])))
			}
		}
	}
	fmt.Println(grid)
	return grid[len(grid)-1][len(grid[0])-1]
}

func main() {
	var grid = [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum(grid))
}
