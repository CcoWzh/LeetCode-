package main

import (
	"fmt"
)

/**
在整个排序过程如图所示，以 arr = [ 8, 1, 4, 6, 2, 3, 5, 7] 为例，
它会把 arr 分成两组 A = [ 8 ] 和 B = [ 1, 4, 6, 2, 3, 5, 7] ，
逐步遍历 B 中元素插入到 A 中，最终构成一个有序序列.
 */
func insertionSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	////这样编程不是最好最简洁的
	//for i := 1; i < len(nums); i++ {
	//	preindx := i - 1
	//	current := nums[i]
	//	//不是nums[preindx] > nums[preindx+1]
	//	for nums[preindx] > current && preindx >= 0 {
	//		nums[preindx+1] = nums[preindx]
	//		preindx--
	//		if preindx < 0 {
	//			break
	//		}
	//	}
	//	//preindx--,所以这里要加回来
	//	nums[preindx+1] = current
	//}

	//这样编程比较简洁
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
		fmt.Println(nums)
	}

	fmt.Println(nums)
	return nums
}

func main() {
	nums := []int{8, 1, 4, 6, 2, 3, 5, 7}
	insertionSort(nums)
}
