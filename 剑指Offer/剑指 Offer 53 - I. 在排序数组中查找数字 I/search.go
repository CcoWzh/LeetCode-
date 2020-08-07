package main

import "fmt"

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	left, right := 0, n
	//寻找左边界
	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}
	//统计次数
	res := 0
	for i := left; i < n; i++ {
		if nums[i] != target {
			break
		}
		res++
	}
	fmt.Println(left, right)
	return res
}

func main() {
	nums := []int{8, 8, 8, 8, 8, 8}
	target := 5
	fmt.Println(search(nums, target))
}
