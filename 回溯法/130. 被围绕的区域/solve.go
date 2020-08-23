package main

import "fmt"

//直接从边缘上的'O'开始遍历即可
func solve(board [][]byte) {
	n := len(board)
	if n == 0 {
		return
	}
	m := len(board[0])
	//预处理
	mark := make([][]int, n)
	for i := 0; i < n; i++ {
		mark[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if board[i][j] == 'O' && (i == 0 || i == n-1 || j == 0 || j == m-1) {
				mark[i][j] = 1
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mark[i][j] == 1 {
				bfs(board, n, m, i, j, mark)
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if mark[i][j] != 1 {
				board[i][j] = 'X'
			}
		}
	}

	fmt.Println(board)
}

//广度优先搜索
func bfs(board [][]byte, n, m, i, j int, mark [][]int) {
	queue := make([][]int, 0)
	queue = append(queue, []int{i, j})

	x := []int{1, -1, 0, 0}
	y := []int{0, 0, 1, -1}
	for len(queue) != 0 {
		curI, curJ := queue[0][0], queue[0][1]
		queue = queue[1:]
		for i := 0; i < 4; i++ {
			newI, newJ := curI+x[i], curJ+y[i]
			if newI < 0 || newI >= n || newJ < 0 || newJ >= m || mark[newI][newJ] == 1 || board[newI][newJ] == 'X' {
				continue
			}
			mark[newI][newJ] = 1
			queue = append(queue, []int{newI, newJ})
		}
	}
}

func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
		{'X', 'O', 'X', 'X'}}
	solve(board)
}
