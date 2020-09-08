package main

import "fmt"

func findErrorNums(nums []int) []int {
	a, b := 0, 0 //a:重复的数字，b:缺失的数字
	n := len(nums)
	for i := 0; i < n; i++ {
		index := abs(nums[i]) - 1 //1到n的整数,不会越界
		if nums[index] < 0 {
			a = index + 1
		} else {
			//注意，这里是由于题目限制，才敢使用相反数标记以及遍历过了
			//不然，这个也得加绝对值后变负数
			nums[index] *= -1
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			b = i + 1
		}
	}
	return []int{a, b}
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func main() {
	nums := []int{1, 2, 2, 4}
	fmt.Println(findErrorNums(nums))
}
