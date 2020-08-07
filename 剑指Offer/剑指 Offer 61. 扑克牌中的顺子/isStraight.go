package main

import (
	"fmt"
	"sort"
)

/**
限制：
数组长度为 5
数组的数取值为 [0, 13] .
 */
func isStraight(nums []int) bool {
	sort.Ints(nums) //排序
	count := 0      //统计有几个0
	for i := 0; i < 5; i++ {
		if nums[i] == 0 {
			count++
		} else if nums[i] != 0 && i+1 < 5 && nums[i] == nums[i+1] {
			//当数组中，除0外，出现对子，就不是顺子
			return false
		}
	}

	for i := 0; i < 4; i++ {
		if nums[i] == 0 { //跳过0
			continue
		}
		count = count - (nums[i+1] - nums[i] - 1)
	}

	return count >= 0
}

func main() {
	nums := []int{0, 0, 1, 2, 3}
	fmt.Println(isStraight(nums))
}
