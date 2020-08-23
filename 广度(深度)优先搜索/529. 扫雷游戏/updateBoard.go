package main

import "fmt"

func updateBoard(board [][]byte, click []int) [][]byte {
	n := len(board)
	if n == 0 {
		return nil
	}
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}
	m := len(board[0])
	//广度优先搜索
	queue := make([][]int, 0)
	queue = append(queue, click)
	dx := []int{1, -1, 0, 0, 1, 1, -1, -1} //上，下，左，右，和所有4个对角线
	dy := []int{0, 0, 1, -1, 1, -1, 1, -1}
	mark := make([][]int, n)
	for i := 0; i < n; i++ {
		mark[i] = make([]int, m)
	}

	for len(queue) != 0 {
		curX, curY := queue[0][0], queue[0][1]
		queue = queue[1:]
		count := 0
		for i := 0; i < 8; i++ {
			newX, newY := curX+dx[i], curY+dy[i]
			if newX < 0 || newX >= n || newY < 0 || newY >= m {
				continue
			}
			if board[newX][newY] == 'M' {
				count++
			}
			//不能在这里入栈，需要当前坐标(curX, curY)改变值以后，才能入队列
		}

		if count > 0 {
			board[curX][curY] = byte(count + '0')
		} else {
			board[curX][curY] = 'B'
			//广度优先搜索
			for i := 0; i < 8; i++ {
				newX, newY := curX+dx[i], curY+dy[i]
				if newX < 0 || newX >= n || newY < 0 || newY >= m || board[newX][newY] != 'E' || mark[newX][newY] == 1 {
					continue
				}
				queue = append(queue, []int{newX, newY})
				mark[newX][newY] = 1 //标记，已经遍历过了
			}
		}
	}

	return board
}

func main() {
	board := [][]byte{
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'M', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'}}
	click := []int{3, 0}
	fmt.Println(updateBoard(board, click))
}
