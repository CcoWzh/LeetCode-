package main

import (
	"fmt"
)

func intersect(nums1 []int, nums2 []int) []int {
	result := make([]int, 0)
	m := make(map[int]bool)

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] && !m[j] {
				result = append(result, nums1[i])
				m[j] = true
				break
			}
		}
	}
	return result
}

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersect(nums1, nums2))
}

//另一个更好的方法:
//两个指针分别指向两个数组的头部
