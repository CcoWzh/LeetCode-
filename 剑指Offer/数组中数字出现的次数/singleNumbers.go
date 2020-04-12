package main

import "fmt"

/**
一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。
请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
 */
func singleNumbers(nums []int) []int {
	tmp := nums[0]
	for i := 1; i < len(nums); i++ {
		tmp = tmp ^ nums[i]
	}
	fmt.Println(tmp)
	return nums
}

func main() {
	nums := []int{1, 2, 10, 4, 1, 4, 3, 3}
	fmt.Println(singleNumbers(nums))

	fmt.Println(3 ^ 3 ^ 3)
}
