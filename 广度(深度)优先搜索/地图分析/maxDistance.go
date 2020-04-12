package main

import "fmt"

//1162. 地图分析
func maxDistance(grid [][]int) int {
	dx := []int{0, 0, 1, -1}
	dy := []int{1, -1, 0, 0}

	queue := make([][]int, 0)
	m, n := len(grid), len(grid[0])
	//把所有的陆地都入队。
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	hasOcean := false
	var point []int
	for len(queue) != 0 {
		point = queue[0]
		queue = queue[1:]
		// 取出队列的元素，将其四周的海洋入队。
		x, y := point[0], point[1]
		for i := 0; i < 4; i++ {
			newX := x + dx[i]
			newY := y + dy[i]
			if newX < 0 || newX >= m || newY < 0 || newY >= n || grid[newX][newY] != 0 {
				continue
			}
			grid[newX][newY] = grid[x][y] + 1 // 这里我直接修改了原数组，因此就不需要额外的数组来标志是否访问
			hasOcean = true
			queue = append(queue, []int{newX, newY})
		}
	}

	// 没有陆地或者没有海洋，返回-1。
	if point == nil || !hasOcean {
		return -1
	}
	// 返回最后一次遍历到的海洋的距离。
	return grid[point[0]][point[1]] - 1
}

func main() {
	grid := [][]int{
		{1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0}}
	fmt.Println(maxDistance(grid))
}
