package main

import "fmt"

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	n := len(image)
	if n == 0 {
		return nil
	}
	m := len(image[0])
	mark := make([][]int, n)
	for i := 0; i < n; i++ {
		mark[i] = make([]int, m)
	}

	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}

	queue := make([][]int, 0)
	queue = append(queue, []int{sr, sc}) //初始化，入队列
	base := image[sr][sc]

	for len(queue) != 0 {
		curX, curY := queue[0][0], queue[0][1]
		image[curX][curY] = newColor
		queue = queue[1:] //出队列
		for i := 0; i < 4; i++ {
			newX, newY := curX+dx[i], curY+dy[i]
			if newX < 0 || newX >= n || newY < 0 || newY >= m || mark[newX][newY] == 1 || image[newX][newY] != base {
				continue
			}
			queue = append(queue, []int{newX, newY})
			mark[newX][newY] = 1
		}
	}

	return image
}

func main() {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1}}
	sr := 1
	sc := 1
	newColor := 2
	fmt.Println(floodFill(image, sr, sc, newColor))
}
