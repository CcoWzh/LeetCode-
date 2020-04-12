package main

import "fmt"

/**
选择排序的思想是，依次从「无序列表」中找到一个最小的元素放到「有序列表」的最后面。
 */
func seelectSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	for i := 0; i < len(nums); i++ {
		mindex := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[mindex] {
				mindex = j
			}
		}
		data := nums[i]
		nums[i] = nums[mindex]
		nums[mindex] = data
	}
	fmt.Println(nums)
	return nums
}

func main() {
	nums := []int{8, 1, 4, 6, 2, 3, 5, 7}
	//nums := []int{65, 58, 95, 10, 57, 62, 13, 106, 78, 23, 85}
	seelectSort(nums)
}
