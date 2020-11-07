package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	backtrack(nums, []int{}, &res)
	return res
}

func backtrack(nums []int, path []int, res *[][]int) {
	n := len(nums)
	if n == 0 {
		buf := make([]int, len(path))
		copy(buf, path)
		*res = append(*res, buf)
		return
	}

	for i := 0; i < n; i++ {
		cur := nums[i]
		//做选择，从选择列表中剔除当前数字
		path = append(path, cur)
		buf := make([]int, n-1)
		copy(buf, nums[0:i])
		copy(buf[i:], nums[i+1:])
		backtrack(buf, path, res)
		//撤销选择
		path = path[:len(path)-1]
		//去重
		for i+1 <= n-1 && nums[i+1] == cur {
			i++
		}
	}
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))
}
