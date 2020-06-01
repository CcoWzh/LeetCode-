package main

import (
	"fmt"
)

var result [][]int

func array(n, m int) [][]int {

	result = make([][]int, 0)
	track := make([]int, 0)
	backtrack(n, m, track, n)

	fmt.Println(result)
	return result
}

//cur当前可做的选择
func backtrack(n, m int, track []int, cur int) {
	//结束条件
	if len(track) > m || cur < 0 {
		return
	}
	//fmt.Println(track)
	if len(track) == m && cur == 0 {
		fmt.Println(track)
		result = append(result, track)
		//fmt.Println("result is :", result)
		return
	}

	//由于数字至少是1，所以从1开始
	for i := 1; i < n; i++ {
		//减枝
		if cur < m-len(track) || cur-i < 0 {
			break
		}
		temp := cur
		//做选择
		cur = cur - i
		track = append(track, i)
		//递归
		backtrack(n, m, track, cur)
		//撤销选择
		cur = temp
		track = track[:len(track)-1]
	}

}

//计算数字n，拆解成m份的排列组合
func main() {
	array(5, 3)
}
