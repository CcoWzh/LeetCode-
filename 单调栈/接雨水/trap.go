package main

import (
	"fmt"
)

//接雨水
func trap(height []int) int {
	//左边界
	left_max := make([]int, len(height))
	right_max := make([]int, len(height))

	for i := len(height) - 1; i >= 0; i-- {
		max := height[i]
		for j := i; j < len(height); j++ {
			if height[j] >= max {
				max = height[j]
			}
		}
		left_max[i] = max
	}

	for i := 0; i < len(height); i++ {
		max := height[i]
		for j := i; j >= 0; j-- {
			if height[j] >= max {
				max = height[j]
			}
		}
		right_max[i] = max
	}

	var result int
	for i := 0; i < len(height); i++ {
		result += min(left_max[i], right_max[i]) - height[i]
	}

	return result
}

func min(a, b int) int {
	if a > b {
		a = b
	}
	return a
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
}
