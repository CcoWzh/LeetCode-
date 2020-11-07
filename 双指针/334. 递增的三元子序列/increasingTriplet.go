package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	small, mid := math.MaxInt64, math.MaxInt64
	for i := 0; i < n; i++ {
		if nums[i] < small {
			small = nums[i]
		} else if nums[i] > small && nums[i] < mid {
			mid = nums[i]
		} else if nums[i] > mid {
			return true
		}
	}
	return false
}

func main() {
	nums := []int{1, 5, 4, 3, 2, 1}
	fmt.Println(increasingTriplet(nums))
}
