package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	backTrace(0, n, k, 1, []int{}, &res)
	return res
}

func backTrace(cur, n, k int, start int, path []int, res *[][]int) {
	if cur == n && len(path) == k {
		buf := make([]int, len(path))
		copy(buf, path)
		*res = append(*res, path)
		return
	}
	if len(path) > k || cur > n {
		return
	}

	base := make([]int, len(path))
	copy(base, path)
	for i := start; i <= 9; i++ {
		if cur+i > n {
			break
		}
		path = append(path, i)
		backTrace(cur+i, n, k, i+1, path, res)
		path = base
	}
}

func main() {
	k := 8
	n := 36
	fmt.Println(combinationSum3(k, n))
}
