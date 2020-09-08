package main

import "fmt"

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	m := make(map[int]int)
	//计算节点的入度值
	for i := 0; i < len(edges); i++ {
		m[edges[i][1]]++
	}
	//统计入度为0的节点
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		if m[i] == 0 {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	n := 6
	edges := [][]int{
		{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}}
	fmt.Println(findSmallestSetOfVertices(n, edges))
}
