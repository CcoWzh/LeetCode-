package main

import (
	"fmt"
)

//先将这个变为图结构
//再判断这个图是否是二分图
func possibleBipartition(N int, dislikes [][]int) bool {
	graph := make([][]int, N)
	for i := 0; i < len(dislikes); i++ {
		a := dislikes[i][0] - 1
		b := dislikes[i][1] - 1
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}
	fmt.Println(graph)

	return isBipartite(graph)
}

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
	N := 10
	//dislikes := [][]int{
	//	{1, 2}, {1, 3}, {2, 4},
	//}
	dislikes := [][]int{
		{4, 7}, {4, 8}, {5, 6}, {1, 6}, {3, 7}, {2, 5}, {5, 8}, {1, 2}, {4, 9}, {6, 10}, {8, 10}, {3, 6}, {2, 10}, {9, 10}, {3, 9}, {2, 3}, {1, 9}, {4, 6}, {5, 7}, {3, 8}, {1, 8}, {1, 7}, {2, 4},
	}

	fmt.Println(possibleBipartition(N, dislikes))
}
