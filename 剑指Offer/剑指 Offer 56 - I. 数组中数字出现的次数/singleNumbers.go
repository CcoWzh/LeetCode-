package main

import "fmt"

func singleNumbers(nums []int) []int {
	n := len(nums)
	ruler := 0
	//异或
	for i := 0; i < n; i++ {
		ruler ^= nums[i]
	}
	//寻找数字的第一个1出现的位置
	index := 0
	for ruler&1 == 0 && ruler != 0 {
		ruler = ruler >> 1
		index++
	}

	a, b := 0, 0
	for i := 0; i < n; i++ {
		//以”数字num的第i位是否为1“为依据，将数组分为两部分
		//每部分，都只包含一个数字，只出现一次
		if isBit(index, nums[i]) {
			a ^= nums[i]
		} else {
			b ^= nums[i]
		}
	}

	return []int{a, b}
}

//判断数字num的第i位是否为1
func isBit(index int, num int) bool {
	num = num >> uint(index)
	return num&1 == 1
}

func main() {
	nums := []int{4, 4, 3, 5}
	fmt.Println(singleNumbers(nums))
}
