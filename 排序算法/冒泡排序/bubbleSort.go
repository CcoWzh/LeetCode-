package main

import "fmt"

//冒泡排序（升序排列）
/**
冒泡排序是通过比较两个相邻元素的大小实现排序，
如果前一个元素大于后一个元素，就交换这两个元素。
这样就会让每一趟冒泡都能找到最大一个元素并放到最后。
 */
func bubbleSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	for i := 0; i < len(nums); i++ {
		//ok := false
		for j := 0; j < len(nums)-i-1; j++ {
			data := nums[j+1]
			if nums[j] > nums[j+1] {
				nums[j+1] = nums[j]
				nums[j] = data
				//ok = true
			}
		}
	}

	fmt.Println(nums)
	return nums
}

func main() {
	nums := []int{8, 1, 4, 6, 2, 3, 5, 7}
	bubbleSort(nums)
}
