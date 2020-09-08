package main

import "fmt"

func kLengthApart(nums []int, k int) bool {
	n := len(nums)
	index, count := 0, 0
	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			index = i + 1
			break
		}
	}

	for index < n {
		if nums[index] == 0 {
			count++
			index++
		} else {
			if count < k {
				return false
			}
			count = 0
			index++
		}

	}

	return true
}

func main() {
	nums := []int{0, 0, 1, 1, 1}
	k := 1
	fmt.Println(kLengthApart(nums, k))
}
