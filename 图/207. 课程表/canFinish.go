package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	n := len(prerequisites)
	if n == 0 {
		return true
	}
	//先将先修课程表变成一个邻接表
	graph := make([][]int, numCourses)
	//计算各个节点的入度情况
	inDegree := make([]int, numCourses)
	for i := 0; i < n; i++ {
		graph[prerequisites[i][0]] = append(graph[prerequisites[i][0]], prerequisites[i][1])
		inDegree[prerequisites[i][1]]++
	}
	fmt.Println(graph, inDegree)

	//广度优先搜索
	queue := make([]int, 0)
	//初始化，先将入度为0的全部加入队列中
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	//开始搜索，如果形成环的话，则res中根本就不能录入这些节点
	res := make([]int, 0)
	for len(queue) > 0 {
		node := queue[0]
		res = append(res, node)
		queue = queue[1:]
		//出队列，同时将这个节点关联的其他节点的入度减1
		for j := 0; j < len(graph[node]); j++ {
			in := graph[node][j]
			inDegree[in]--
			if inDegree[in] == 0 {
				queue = append(queue, in)
			}
		}
	}
	fmt.Println(res)

	return len(res) == numCourses
}

func main() {
	numCourses := 6
	prerequisites := [][]int{
		{0, 2}, {0, 1}, {1, 2}, {0, 3}, {5, 3}, {5, 4}}
	fmt.Println(canFinish(numCourses, prerequisites))
}
