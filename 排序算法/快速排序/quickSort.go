package main

import (
	"fmt"
)

/**
快速排序：指针交换（挖坑填数）+分治法
 */
//func quickSort(nums []int, left, right int) []int {
//	//temp := nums[left]
//	temp := nums[(left+right)/2]
//	i, j := left, right
//	//判断条件时 i <= j，不是<
//	for i <= j {
//		//向左移动，直到找到大于支点的元素
//		for nums[i] < temp {
//			i++
//		}
//		//向右移动，直到找到小于支点的元素
//		for nums[j] > temp {
//			j--
//		}
//		//交换两个元素，让左边的大于支点，右边的小于支点
//		if i <= j {
//			nums[i], nums[j] = nums[j], nums[i]
//			i++
//			j--
//		}
//	}
//	//递归左边，进行快速排序
//	if left < j {
//		quickSort(nums, left, j)
//	}
//	//递归右边，进行快速排序
//	if right > i {
//		quickSort(nums, i, right)
//	}
//
//	return nums
//}

func main() {
	nums := []int{-1, -2, 8, 1, 1, 1, 2, 3, 5, 7}
	//nums := []int{65, 58, 95, 10, 57, 62, 13, 106, 78, 23, 85}
	sort := QuickSort(nums, 0, len(nums)-1)
	fmt.Println(sort)
}

//快排的另一种写法
func QuickSort(nums []int, left, right int) []int {
	if left < right {
		index := Paritition(nums, left, right)
		QuickSort(nums, left, index-1)
		QuickSort(nums, index+1, right)
	}

	return nums
}

func Paritition(nums []int, low, height int) int {
	pivot := nums[low]
	for low < height {
		for low < height && nums[height] >= pivot {
			height--
		}
		nums[low] = nums[height]
		for low < height && nums[low] <= pivot {
			low++
		}
		nums[height] = nums[low]
	}
	nums[low] = pivot

	return low
}
