package main

import "fmt"

//染色法，但是题目没说一定是连通图
//这点是个坑
func isBipartite(graph [][]int) bool {
	color := make([]int, len(graph))

	for i := 0; i < len(graph); i++ {
		if color[i] != 0 {
			continue
		}
		queue := make([]int, 0)
		queue = append(queue, i)
		color[i] = 1

		for len(queue) != 0 {
			index := queue[0]
			queue = queue[1:] //出队列
			n := len(graph[index])
			for i := 0; i < n; i++ {
				curIndex := graph[index][i]
				if color[curIndex] == 0 {
					color[curIndex] = 0 - color[index]
					queue = append(queue, curIndex)
				} else {
					if color[curIndex] == color[index] {
						return false
					}
				}
			}
		}
	}
	fmt.Println(color)

	return true
}

func main() {
	graph := [][]int{
		{}, {3}, {}, {1},
	}
	fmt.Println(isBipartite(graph))
}
