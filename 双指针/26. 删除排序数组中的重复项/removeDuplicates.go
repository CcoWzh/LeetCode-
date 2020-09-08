package main

import "fmt"

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	fast, low := 0, 0
	for fast < len(nums) {
		if nums[fast] == nums[low] {
			fast++
		} else {
			low++ //注意，这个low一定要再前面执行
			nums[low] = nums[fast]
		}
	}
	fmt.Println(nums[:low+1])
	fmt.Println(nums)
	return low + 1
}

//给定一个排序数组
func main() {
	nums := []int{1}
	removeDuplicates(nums)
}
