package main

import "fmt"

//其中nums1 是 nums2 的子集
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	//单调递减栈
	stack := make([]int, 0)
	numMap := make(map[int]int, 0)
	n, m := len(nums2), len(nums1)
	//if m > n || n < 0 { //题目判断时，似乎没有要求满足“nums1 是 nums2 的子集”这一条件
	//	return nil
	//}
	//单调栈
	for i := 0; i < n; i++ {
		temp := nums2[i]
		for len(stack) != 0 && temp > stack[len(stack)-1] {
			v := stack[len(stack)-1]
			numMap[v] = temp
			//出栈
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, temp)
	}
	//fmt.Println(numMap)

	result := make([]int, m)
	for i := 0; i < m; i++ {
		temp := nums1[i]
		if j, ok := numMap[temp]; ok {
			result[i] = j
		} else {
			result[i] = -1
		}
	}

	return result
}

func main() {
	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}
	fmt.Println(nextGreaterElement(nums1, nums2))
}
