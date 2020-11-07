package main

import "fmt"

func search(nums []int, target int) bool {
	n := len(nums)
	if n == 0 {
		return false
	} else if n == 1 {
		return nums[0] == target
	}
	left, right := 0, n-1

	for left <= right {
		mid := (left + right) / 2
		cur := nums[mid]
		if cur == target {
			return true
		}
		if nums[left] == cur {
			left++
			continue
		}
		//以nums[mid]为界限，总有一半是有序的
		if nums[left] < cur { //说明前半部分有序
			//如果target在这个前半部分，则将范围缩小到这里
			if nums[left] <= target && target < cur {
				right = mid - 1
			} else {
				//否者，就去另一半去查找
				left = mid + 1
			}
		} else { //后半部分有序
			if cur < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}

func main() {
	nums := []int{1, 3, 5}
	target := 1
	fmt.Println(search(nums, target))
}
