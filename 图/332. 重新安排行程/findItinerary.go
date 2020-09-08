package main

import (
	"fmt"
	"sort"
)

func findItinerary(tickets [][]string) []string {
	//首先，需要构建邻接图
	n := len(tickets)
	graph := make(map[string][]string)
	for i := 0; i < n; i++ {
		from, to := tickets[i][0], tickets[i][1]
		graph[from] = append(graph[from], to)
	}
	//对邻接图中的值按照字典序排序
	for _, v := range graph {
		sort.Strings(v)
	}
	res := make([]string, 0)
	dfs("JFK", graph, &res)
	//逆序输出
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}

	return res
}

//Hierholzer算法
//graph是map类型，属于浅拷贝
func dfs(curStr string, graph map[string][]string, res *[]string) {
	for {
		if v, ok := graph[curStr]; !ok || len(v) == 0 {
			break
		}
		nextStr := graph[curStr][0]       //取出第一个节点
		graph[curStr] = graph[curStr][1:] //删除这条边
		dfs(nextStr, graph, res)
	}
	*res = append(*res, curStr)
}

func main() {
	tickets := [][]string{
		{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}
	fmt.Println(findItinerary(tickets))

}
