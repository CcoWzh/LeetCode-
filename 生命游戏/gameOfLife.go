package main

import "fmt"

func gameOfLife(board [][]int) {
	temp := make([][]int, len(board))

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			count := countLive(board, i, j)
			if (count < 2 || count > 3) && board[i][j] == 1 {
				temp[i] = append(temp[i], 0)
			} else if (count == 2 || count == 3) && board[i][j] == 1 {
				temp[i] = append(temp[i], 1)
			} else if count == 3 && board[i][j] == 0 {
				temp[i] = append(temp[i], 1)
			} else {
				temp[i] = append(temp[i], board[i][j])
			}
		}
	}

	copy(board, temp)

	fmt.Println(board)

}

//统计一个数字上下左右有多少个1
func countLive(board [][]int, m, n int) int {
	x := []int{0, 0, 1, -1, 1, -1, 1, -1}
	y := []int{1, -1, 0, 0, 1, 1, -1, -1}

	count := 0
	for i := 0; i < 8; i++ {
		index_x, index_y := m+x[i], n+y[i]
		if index_x < 0 || index_y < 0 || index_x >= len(board) || index_y >= len(board[0]) {
			continue
		}
		if board[index_x][index_y] == 1 {
			count++
		}
	}
	return count
}

func main() {
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0}}
	gameOfLife(board)
}
