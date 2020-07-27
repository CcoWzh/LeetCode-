package main

import "fmt"

func findErrorNums(nums []int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	a, b := 0, 0
	for i := 1; i <= len(nums); i++ {
		if m[i] == 2 {
			a = i
		} else if m[i] == 0 {
			b = i
		}
	}

	return []int{a, b}
}

func main() {
	nums := []int{1, 1}
	fmt.Println(findErrorNums(nums))
}
