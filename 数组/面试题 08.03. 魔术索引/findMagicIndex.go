package main

import "fmt"

func findMagicIndex(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == i {
			return i
		}
	}
	return -1
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 2, 6, 7, 8, 9}
	fmt.Println(findMagicIndex(nums))
}
