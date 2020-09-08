package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	nums := make([]int, 0)

	//x必是大于1的，想啥呢
	for x > 0 {
		n := x % 10
		nums = append(nums, n)
		x = x / 10
	}

	for l, r := 0, len(nums)-1; l <= r; {
		if nums[l] == nums[r] {
			l++
			r--
		} else {
			return false
		}
	}

	fmt.Println(nums)

	return true
}

func main() {
	x := -0
	fmt.Println(isPalindrome(x))
}
