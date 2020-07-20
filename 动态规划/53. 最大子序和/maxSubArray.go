package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return -2147483648
	}
	max, n := nums[0], len(nums)

	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			if max < sum {
				max = sum
			}
		}
	}
	return max
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
