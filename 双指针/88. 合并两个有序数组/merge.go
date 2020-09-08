package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	k := n + m - 1
	p1, p2 := m-1, n-1 //两个指针
	for p1 >= 0 || p2 >= 0 {
		a, b := -1000000, -100000
		if p1 >= 0 {
			a = nums1[p1]
		}
		if p2 >= 0 {
			b = nums2[p2]
		}
		if a > b {
			nums1[k] = a
			p1--
		} else {
			nums1[k] = b
			p2--
		}
		k--
	}

	fmt.Println(nums1)
}

func main() {
	m, n := 3, 3
	nums1 := make([]int,0,n+m)
	nums1 = []int{1, 2, 3}
	nums2 := []int{2, 5, 6}
	merge(nums1, m, nums2, n)
}
