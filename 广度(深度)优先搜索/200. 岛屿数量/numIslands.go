package main

import "fmt"

func numIslands(grid [][]byte) int {
	n := len(grid)
	if n == 0 {
		return 0
	}
	m := len(grid[0])
	res := 0
	queue := make([][]int, 0)
	dpX := []int{1, -1, 0, 0}
	dpY := []int{0, 0, 1, -1}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				res++ //更新结果
				queue = append(queue, []int{i, j})
				for len(queue) != 0 {
					x, y := queue[0][0], queue[0][1]
					queue = queue[1:]
					for k := 0; k < 4; k++ {
						newX, newY := x+dpX[k], y+dpY[k]
						if newX < 0 || newX >= n || newY < 0 || newY >= m || grid[newX][newY] != '1' {
							continue
						}
						grid[newX][newY] = '2'
						queue = append(queue, []int{newX, newY})
					}
				}
			}
		}
	}
	return res
}

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'}}
	fmt.Println(numIslands(grid))
}
