package main

import "fmt"

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i:=0;i<n;i++{
		res[i] = make([]int,n)
	}
	base := n / 2
	if base*2 < n {
		base += 1
	}

	num := 1
	for i := 0; i < base; i++ {
		end := n - i
		start := i
		//第1步，行，左到右
		for j := start; j < end; j++ {
			res[start][j] = num
			num++
		}
		//第2步，列，上到下
		for j := start + 1; j < end; j++ {
			res[j][end-1] = num
			num++
		}
		//第3步，行，右到左
		for j := end - 2; j >= start; j-- {
			res[end-1][j] = num
			num++
		}
		//第4步，列，下到上
		for j := end - 2; j > start; j-- {
			res[j][start] = num
			num++
		}
	}
	fmt.Println(res)
	return res
}

func main() {
	n := 3
	generateMatrix(n)
}
