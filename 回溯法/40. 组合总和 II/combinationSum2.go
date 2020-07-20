package main

import (
	"fmt"
	"sort"
)

var result [][]int

func combinationSum2(candidates []int, target int) [][]int {
	result = make([][]int, 0)
	sort.Ints(candidates)
	backtrack(candidates, []int{}, 0, target, 0)

	return result
}

//selectList可选择的列表
//tarck已经选择的路径, curSum 当前和, target 目标, index 当前下标
//如果要加result [][]int，需要加指针，不然会报错
func backtrack(selectList []int, tarck []int, curSum int, target int, index int) {
	if curSum == target {
		//这里一定要copy，要不然结果result会变，这是由于track会变的缘故，浅拷贝
		s := make([]int, len(tarck))
		copy(s, tarck)
		result = append(result, s)
		//fmt.Println("result is ", result)
	}

	for i := index; i < len(selectList); i++ {
		temp := selectList[i]
		if curSum+temp > target {
			continue
		}
		tarck = append(tarck, selectList[i])
		backtrack(delete(selectList, i), tarck, curSum+temp, target, i)
		tarck = tarck[:len(tarck)-1]
		//除重复,如：1,1,1,5,6；去除1,1,1
		for i+1 <= len(selectList)-1 && selectList[i+1] == temp {
			i++
		}
	}
}

func delete(nums []int, i int) []int {
	if i >= len(nums) {
		return nil
	}
	result1 := make([]int, len(nums))
	copy(result1, nums)
	temp := result1[i+1:]
	result1 = result1[:i]
	result1 = append(result1, temp...)
	return result1
}

func main() {
	candidates := []int{3, 1, 3, 5, 1, 1}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
	//fmt.Println(delete(candidates, 0),candidates)
}
