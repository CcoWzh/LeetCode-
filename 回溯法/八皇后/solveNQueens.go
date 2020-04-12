package main

import "fmt"

//这里声明一个全局变量用来存储所有的排列
var result [][]string

/**
待完成
 */
func solveNQueens(n int) [][]string {
	//每次调用重置result全局变量，防止结果缓存
	result = make([][]string, 0)
	track := make([]string, 0)
	backtrack(track, 0)
	return result
}

func backtrack(track []string, row int) {

}

func main() {
	fmt.Println(solveNQueens(4))
}
