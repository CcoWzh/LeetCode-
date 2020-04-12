package main

import "fmt"

func surfaceArea(grid [][]int) int {
	var sum int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			//一定要加判断条件grid[i][j] > 0
			if grid[i][j] > 0 {
				sum += 4*grid[i][j] + 2
				if i > 0 {
					sum -= min(grid[i-1][j], grid[i][j]) << 1
				}
				if j > 0 {
					sum -= min(grid[i][j-1], grid[i][j]) << 1
				}
				//四面相减
				//if i > 0 {
				//	sum -= min(grid[i-1][j], grid[i][j])
				//}
				//
				//if i+1 < len(grid) {
				//	sum -= min(grid[i+1][j], grid[i][j])
				//}
				//
				//if j+1 < len(grid[i]) {
				//	sum -= min(grid[i][j+1], grid[i][j])
				//}
				//
				//if j > 0 {
				//	sum -= min(grid[i][j-1], grid[i][j])
				//}
			}

		}
	}
	return sum
}

func min(a, b int) int {
	if a > b {
		a = b
	}
	return a
}

func main() {
	grid := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1}}
	fmt.Println(surfaceArea(grid))
}
