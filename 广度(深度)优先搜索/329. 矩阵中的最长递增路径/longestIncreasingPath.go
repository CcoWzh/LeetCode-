package main

import "fmt"

var x = []int{1, -1, 0, 0}
var y = []int{0, 0, 1, -1}

func longestIncreasingPath(matrix [][]int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])
	//记忆化搜索
	memory := make([][]int, n)
	for i := 0; i < n; i++ {
		memory[i] = make([]int, m)
	}
	M := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			M = max(M, dfs(matrix, memory, i, j))
		}
	}
	fmt.Println(memory)
	return M
}

//返回从(第row行,column列)出发，所能到达的最大递增路径
func dfs(matrix, memory [][]int, row, column int) int {
	if memory[row][column] != 0 {
		return memory[row][column]
	}
	n, m := len(matrix), len(matrix[0])

	for i := 0; i < 4; i++ {
		newX, newY := row+x[i], column+y[i]
		if newX >= 0 && newX < n && newY >= 0 && newY < m && matrix[newX][newY] > matrix[row][column] {
			memory[row][column] = max(memory[row][column], dfs(matrix, memory, newX, newY))
		}
	}
	memory[row][column]++

	return memory[row][column]
}

func max(a, b int) int {
	if a < b {
		a = b
	}
	return a
}

func main() {
	//matrix := [][]int{
	//	{9, 9, 4},
	//	{6, 6, 8},
	//	{2, 1, 1}}
	matrix := [][]int{{}}

	fmt.Println(longestIncreasingPath(matrix))
}
