package main

import "fmt"

func canJump(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return true
	}
	maxIndex := 0
	for i := 0; i < n-1; i++ {
		maxIndex = max(maxIndex, i+nums[i]) //表示当前可以跳跃到的最远的下标
		if maxIndex <= i { //nums[i] == 0
			return false
		}
	}
	return maxIndex >= n-1
}

func max(a, b int) int {
	if a < b {
		a = b
	}
	return a
}

func main() {
	nums := []int{0}
	fmt.Println(canJump(nums))
}
