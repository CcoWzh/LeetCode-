package main

import "fmt"

func findDuplicates(nums []int) []int {
	n := len(nums)
	res := make([]int, 0)
	if n == 0 {
		return res
	}

	for i := 0; i < n; i++ {
		for nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			res = append(res, nums[i])
		}
	}
	return res
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(nums))
}
