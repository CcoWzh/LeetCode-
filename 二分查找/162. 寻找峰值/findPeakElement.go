package main

import "fmt"

//这题居然也能使用二分法？
func findPeakElement(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)/2
		if mid+1 < n && nums[mid] > nums[mid+1] { //说明在降序
			right = mid
		} else {
			//升序阶段
			left = mid + 1
		}
	}
	return left
}

func main() {
	nums := []int{1, 2, 1, 3, 5, 6, 4}
	fmt.Println(findPeakElement(nums))
}
