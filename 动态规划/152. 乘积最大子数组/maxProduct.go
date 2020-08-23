package main

import "fmt"

func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dpMax := make([]int, n)
	dpMin := make([]int, n)
	dpMax[0], dpMin[0] = nums[0], nums[0]

	res := dpMax[0]
	for i := 1; i < n; i++ {
		dpMax[i] = max(nums[i], max(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i]))
		dpMin[i] = min(nums[i], min(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i]))
		if res < dpMax[i] {
			res = dpMax[i]
		}
	}

	//fmt.Println(dpMax, dpMin)
	return res
}

func max(a, b int) int {
	if a < b {
		a = b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		a = b
	}
	return a
}

func main() {
	nums := []int{2, 3, -2, 4}
	fmt.Println(maxProduct(nums))
}
