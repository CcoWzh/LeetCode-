package main

import "fmt"

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	mark := make([]int, n)
	//广度优先遍历
	queue := make([]int, 0)
	queue = append(queue, 0)

	for len(queue) != 0 {
		index := queue[0]
		queue = queue[1:]
		mark[index] = 1
		for i := 0; i < len(rooms[index]); i++ {
			cur := rooms[index][i]
			if mark[cur] != 1 {
				queue = append(queue, cur)
			}
		}
	}

	for i := 0; i < n; i++ {
		if mark[i] == 0 {
			return false
		}
	}
	return true
}

func main() {
	rooms := [][]int{{1, 3}, {3, 0, 1}, {2}, {0}}
	fmt.Println(canVisitAllRooms(rooms))
}
