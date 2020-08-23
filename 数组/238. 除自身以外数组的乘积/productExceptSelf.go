package main

import "fmt"

/**
238. 除自身以外数组的乘积
 */
func productExceptSelf(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	//前缀+后缀
	pre := make([]int, n)
	suf := make([]int, n)
	result := make([]int, n)
	pre[0], suf[n-1] = 1, 1

	for i := 1; i < n; i++ {
		pre[i] = nums[i-1] * pre[i-1]
	}

	for j := n - 2; j >= 0; j-- {
		suf[j] = nums[j+1] * suf[j+1]
	}

	for i := 0; i < n; i++ {
		result[i] = pre[i] * suf[i]
	}

	return result
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}
