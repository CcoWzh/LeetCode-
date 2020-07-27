package main

import "fmt"

//你可以假设数组中不存在重复元素
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) / 2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

func main() {
	nums := []int{9, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(findMin(nums))
}
