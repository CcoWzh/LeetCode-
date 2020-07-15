package main

import (
	"fmt"
)

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 {
		return []int{0}
	} else if k == len(arr) {
		return arr
	}
	partitionArray(arr, 0, len(arr)-1, k)
	return arr[:k]
}

func quickSort(nums []int, left, right int) int {
	temp := nums[left]
	//temp := nums[(left+right)/2]
	i, j := left, right
	//判断条件时 i <= j，不是<
	for i <= j {
		//向左移动，直到找到大于支点的元素
		for nums[i] < temp {
			i++
		}
		//向右移动，直到找到小于支点的元素
		for nums[j] > temp {
			j--
		}
		if i >= j {
			break
		}
		//交换两个元素，让左边的大于支点，右边的小于支点
		if i <= j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i], nums[j] = nums[j], nums[i]
	return j
}

func partitionArray(arr []int, lo, hi, k int) {
	// 做一次 partition 操作
	m := quickSort(arr, lo, hi)
	// 此时数组前 m 个数，就是最小的 m 个数
	if k == m {
		// 正好找到最小的 k(m) 个数
		return
	} else if k < m {
		// 最小的 k 个数一定在前 m 个数中，递归划分
		partitionArray(arr, lo, m-1, k)
	} else {
		// 在右侧数组中寻找最小的 k-m 个数
		partitionArray(arr, m+1, hi, k)
	}
}

func main() {
	arr := []int{0, 1, 2, 3}
	k := 3
	fmt.Println(getLeastNumbers(arr, k))
}
