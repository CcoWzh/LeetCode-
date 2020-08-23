package main

import "fmt"

/**
迪杰斯特拉最短路径算法,可以计算从节点k到所以其可达节点之间的最短路径
 */
func networkDelayTime(times [][]int, N int, K int) int {
	var MAXWEIGHT = 100000000
	//转换成邻接图
	n := len(times)
	if n == 0 {
		return 0
	}
	graph := make(map[int][][]int)
	for i := 0; i < n; i++ {
		cur := times[i][0]
		graph[cur] = append(graph[cur], []int{times[i][1], times[i][2]}) //节点+权重
	}
	//迪杰斯特拉最短路径算法
	distance := make(map[int]int) //顶点k到任意其他顶点的距离
	for i := 1; i <= N; i++ {
		if i == K {
			distance[K] = 0
		} else {
			distance[i] = MAXWEIGHT
		}
	}
	T := make([]bool, N+1) //记录遍历过的顶点集T

	for {
		curNode := -1
		curDist := MAXWEIGHT
		for i := 1; i <= N; i++ { //遍历distance，找到距离k最近的顶点，且这个节点没有被访问过（是一个估计值，而不是确定值）
			if !T[i] && distance[i] < curDist {
				curDist = distance[i]
				curNode = i
			}
		}
		if curNode < 0 { //如果curNode没有改变，则说明从节点k到其他可达的节点都遍历完了
			break
		}
		T[curNode] = true //否则，将其标记为以及遍历了，即，是确定值
		//寻找和当前节点连接的其他顶点，更新他们的距离
		list := graph[curNode]
		for i := 0; i < len(list); i++ {
			nextNode := list[i][0]
			distance[nextNode] = min(distance[nextNode], curDist+list[i][1])
		}
	}

	//fmt.Println(distance)
	res := 0
	for i := 1; i <= N; i++ {
		if distance[i] == MAXWEIGHT { //如果距离中还存在有MAXWEIGHT，说明节点k到该节点不可达
			return -1
		}
		res = max(res, distance[i])
	}

	return res
}

func min(a, b int) int {
	if a > b {
		a = b
	}
	return a
}

func max(a, b int) int {
	if a < b {
		a = b
	}
	return a
}

func main() {
	times := [][]int{
		{2, 1, 1},
		{2, 3, 1},
		{3, 4, 1}}
	N := 4
	K := 2
	fmt.Println(networkDelayTime(times, N, K))
}
