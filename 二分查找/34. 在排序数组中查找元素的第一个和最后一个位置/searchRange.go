package main

import "fmt"

func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}
	left, right := 0, n-1
	//找到左边界
	for left < right {
		mid := left + (right-left)/2
		cur := nums[mid]
		if cur >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	i, j := left, left
	//判断寻找的左边界是否是目标值
	if nums[i] != target {
		return []int{-1, -1}
	}
	for j < n { //寻找右边界
		if nums[j] != nums[i] {
			j--
			break
		}
		if j == n-1 {
			break
		}
		j++
	}

	return []int{i, j}
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 8}
	target := 8
	fmt.Println(searchRange(nums, target))
}
