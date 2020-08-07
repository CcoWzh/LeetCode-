package main

import "fmt"

//归并排序
func mergeSort(arr []int) []int {
	n := len(arr)
	if n < 2 {
		return arr
	}
	mid := n / 2

	return merge(mergeSort(arr[:mid]), mergeSort(arr[mid:]))
}

//合并两个有序数组
/**
1. 申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列；
2. 设定两个指针，最初位置分别为两个已经排序序列的起始位置；
3. 比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置；
4. 重复步骤 3 直到某一指针达到序列尾；
5. 将另一序列剩下的所有元素直接复制到合并序列尾。
 */
func merge(left []int, right []int) []int {
	n, m := len(left), len(right)
	l, r := 0, 0
	res := make([]int, 0)

	for l < n && r < m {
		if left[l] < right[r] {
			res = append(res, left[l])
			l++
		} else {
			res = append(res, right[r])
			r++
		}
	}

	if l < n {
		res = append(res, left[l:]...)
	}
	if r < m {
		res = append(res, right[r:]...)
	}
	return res
}

func main() {
	//fmt.Println(merge([]int{1, 2, 5, 8}, []int{3, 4, 6, 9}))
	fmt.Println(mergeSort([]int{1, 5, 3, 2, 8, 9, 100, 9, 3, 10}))
}
