package main

import "fmt"

func exchange(nums []int) []int {
	left, right := 0, len(nums)-1
	//左右指针
	for left < right {
		//左指针一直右移，直到遇到偶数
		for nums[left]%2 == 1 && left < right {
			left++
		}
		//右指针一直左移，知道遇到奇数停止
		for nums[right]%2 == 0 && right > left {
			right--
		}
		//交换两个数字
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}

	return nums
}

func main() {
	nums := []int{1, 3, 2, 5, 7, 4}
	fmt.Println(exchange(nums))
}
