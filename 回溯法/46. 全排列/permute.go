package main

import (
	"fmt"
)

//这里声明一个全局变量用来存储所有的排列
var result [][]int

//全排列
func permute(nums []int) [][]int {
	//每次调用重置result全局变量，防止结果缓存
	result = make([][]int, 0)
	track := make([]int, 0)
	n := len(nums)
	backtrack(nums, track, n)
	return result
}

func backtrack(nums, track []int, n int) {
	//触发结束条件
	if len(track) == n {
		result = append(result, track)
		//fmt.Println("result is ",result)
		return
	}
	for i := 0; i < len(nums); i++ {
		//做选择
		track := append(track, nums[i])
		buf := make([]int, len(nums)-1)
		copy(buf, nums[0:i])      //----都得先保存当前路径
		copy(buf[i:], nums[i+1:]) // 去除了第i个元素
		//进⼊下⼀层决策树
		//fmt.Println("buf is ", buf)
		//fmt.Println("track is ", track)
		backtrack(buf, track, n)
		//取消选择
		track = track[:len(track)-1]
	}
}

func main() {
	nums := []int{4, 6, 7, 7}
	fmt.Println(permute(nums))
}
