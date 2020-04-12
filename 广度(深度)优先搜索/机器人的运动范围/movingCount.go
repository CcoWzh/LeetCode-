package main

import "fmt"

func movingCount(m int, n int, k int) int {
	dx := []int{0, 1}
	dy := []int{1, 0}
	queue := make([][]int, 0)
	mat := make([][]int, 0)
	for i := 0; i < m; i++ {
		mat = append(mat, make([]int, n))
	}
	mat[0][0] = 1
	queue = append(queue, []int{0, 0})

	result := 1
	for len(queue) != 0 {
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]
		// 取出队列的元素，将其四周的海洋入队。
		for i := 0; i < 2; i++ {
			newX := x + dx[i]
			newY := y + dy[i]
			if newX < 0 || newX > m-1 || newY < 0 || newY > n-1 || mat[newX][newY] == 1 || sum(newX, newY) > k {
				continue
			}
			//fmt.Println(newX, newY, sum(newX, newY) > 0)
			queue = append(queue, []int{newX, newY})
			result ++
			mat[newX][newY] = 1

		}
	}
	return result
}

func sum(a, b int) int {
	sum := 0
	for a%10 != 0 {
		sum += a % 10
		a = a / 10
	}
	for b%10 != 0 {
		sum += b % 10
		b = b / 10
	}
	return sum
}

func main() {
	m, n, k := 35, 15, 9
	fmt.Println(movingCount(m, n, k))
}
