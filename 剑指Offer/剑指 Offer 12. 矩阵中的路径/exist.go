package main

import "fmt"

func exist(board [][]byte, word string) bool {
	n := len(board)
	if n == 0 {
		return false
	}
	m := len(board[0])
	target := word[0]

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == target {
				//标记搜索过的路径
				used := make([][]int, n)
				for k := 0; k < n; k++ {
					used[k] = make([]int, m)
				}
				//BFS搜索
				if dfs(board, used, i, j, "", word) {
					return true
				}
			}
		}
	}

	return false
}

func dfs(board [][]byte, used [][]int, x, y int, path, word string) bool {
	//做判断
	if path == word {
		return true
	}
	l := len(path)
	n, m := len(board), len(board[0])
	if x < 0 || x >= n || y < 0 || y >= m || used[x][y] != 0 || board[x][y] != word[l] {
		return false
	}
	//做选择
	path += string(board[x][y])
	used[x][y] = 1
	//dfs搜索
	res := dfs(board, used, x+1, y, path, word) || dfs(board, used, x-1, y, path, word) || dfs(board, used, x, y+1, path, word) || dfs(board, used, x, y-1, path, word)
	//还原选择
	path = path[:len(path)-1]
	used[x][y] = 0
	return res
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'E', 'S'},
		{'A', 'D', 'E', 'E'}}
	word := "ABCESEEEFS"
	fmt.Println(exist(board, word))
}
