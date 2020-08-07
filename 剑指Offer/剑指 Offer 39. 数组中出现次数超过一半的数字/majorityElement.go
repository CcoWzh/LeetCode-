package main

import "fmt"

func majorityElement(nums []int) int {
	m := make(map[int]int)
	n := len(nums)

	for i := 0; i < n; i++ {
		m[nums[i]]++
		if m[nums[i]] > n/2 {
			return nums[i]
		}
	}
	return 0
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 2, 2}
	fmt.Println(majorityElement(nums))
}
