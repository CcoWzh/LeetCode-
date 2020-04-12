package main

import "fmt"

/** 二进制和位运算
1 << uint(len(nums)) ,即：2的n次方
i >> uint(j) & 1  即：i的二进制表示，用1来逐个遍历,
如101表示选择数组中0号和2号位置
 */
func subsets(nums []int) [][]int {
	result := make([][]int, 0)

	for i := 0; i < (1 << uint(len(nums))); i++ {
		temp := make([]int, 0)
		for j := 0; j < len(nums); j++ {
			if (i >> uint(j) & 1) == 1 {
				temp = append(temp, nums[j])
			}
		}
		result = append(result, temp)
	}
	fmt.Println(result)
	return result
}

func main() {
	nums := []int{1, 2, 3}
	subsets(nums)
}
