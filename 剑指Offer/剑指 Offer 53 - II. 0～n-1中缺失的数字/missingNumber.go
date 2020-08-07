package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums)
	result := 0
	if n <= 0 {
		return result
	}
	result ^= n
	for i := 0; i < n; i++ {
		result ^= i ^ nums[i]
	}

	return result
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(missingNumber(nums))
}
