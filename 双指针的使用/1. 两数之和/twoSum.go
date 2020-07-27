package main

import "fmt"

/**
给定一个整数数组 nums 和一个目标值 target，
请你在该数组中找出和为目标值的那 两个 整数，
并返回他们的数组下标。
 */
func twoSum(nums []int, target int) []int {
	var result []int
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			result = append(result, j)
			result = append(result, i)
		}
		m[v] = i
	}
	return result
}

func main() {
	nums := []int{1, 2, 5, 8, 3}
	fmt.Println(twoSum(nums, 9))
}
