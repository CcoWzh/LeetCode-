package main

import "fmt"

func findUnsortedSubarray(nums []int) int {
	//需要找到从左到右，最后一个下降的数的位置
	//找到从右到左，最后一个上升的数的位置
	n := len(nums)
	start, end := -1, 0
	max, min := nums[0], nums[n-1]
	for i := 0; i < n; i++ {
		//寻找右边界
		if nums[i] >= max { //如果一直是升序的
			max = nums[i]
		} else { //如果有降序的数
			start = i
		}
		//寻找左边界
		if nums[n-1-i] <= min { //如果一直是降序的
			min = nums[n-1-i]
		} else {
			end = n - 1 - i
		}
	}
	//start:右边界, end:左边界
	return start - end + 1
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(findUnsortedSubarray(nums))
}
