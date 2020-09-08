package main

import "fmt"

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	backTrack(1, n, k, 0, []int{}, &res)
	return res
}

//start:搜索的起点, n, k, count:计数，也就睡path的长度
func backTrack(start, n, k, count int, path []int, res *[][]int) {
	if count == k {
		buf := make([]int, k)
		copy(buf, path)
		*res = append(*res, buf)
		return
	}

	base := path
	for i := start; i <= n; i++ { //边界就是 i <= n
		path = append(path, i)
		backTrack(i+1, n, k, count+1, path, res)
		path = base
	}
}

func main() {
	n := 5
	k := 1
	fmt.Println(combine(n, k))
}
