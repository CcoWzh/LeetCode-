package main

import "fmt"

//999. 车的可用捕获量
//这题可以使用方向数组 来简化代码量：
//int dx[4] = {0, 1, 0, -1};
//int dy[4] = {1, 0, -1, 0};
func numRookCaptures(board [][]byte) int {
	result := 0
	index_x, index_y := 8, 8
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if string(board[i][j]) == "R" {
				index_x, index_y = i, j
			}
		}
	}
	if index_y == 8 && index_x == 8 {
		return 0
	}

	for i := index_x; i >= 0; i-- {
		if string(board[i][index_y]) == "B" {
			break
		}
		if string(board[i][index_y]) == "p" {
			result++
			break
		}
	}
	for i := index_x; i < 8; i++ {
		if string(board[i][index_y]) == "B" {
			break
		}
		if string(board[i][index_y]) == "p" {
			result++
			break
		}
	}
	for j := index_y; j < 8; j++ {
		if string(board[index_x][j]) == "B" {
			break
		}
		if string(board[index_x][j]) == "p" {
			result++
			break
		}
	}
	for j := index_y; j >= 0; j-- {
		if string(board[index_x][j]) == "B" {
			break
		}
		if string(board[index_x][j]) == "p" {
			result++
			break
		}
	}
	return result
}

func main() {
	board := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'p', 'p', '.', 'R', '.', 'p', 'B', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'B', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'}}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Print(string(board[i][j]))
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
	fmt.Println(numRookCaptures(board))
}
