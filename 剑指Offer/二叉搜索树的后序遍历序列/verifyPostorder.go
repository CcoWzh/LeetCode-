package main

import "fmt"

func verifyPostorder(postorder []int) bool {
	if len(postorder) < 2 {
		return true
	}
	n := len(postorder)
	root := postorder[n-1]
	index := 0
	for postorder[index] < root && index < n {
		index++
	}
	for i := index; i < n-1; i++ {
		if postorder[i] < root {
			return false
		}
	}
	left := verifyPostorder(postorder[:index])
	right := verifyPostorder(postorder[index : n-1])
	return left && right
}

//输入: [1,3,2,6,5]
//输出: true
func main() {
	postorder := []int{2, 7, 9, 12, 11, 8}
	fmt.Println(verifyPostorder(postorder))

	nums := []int{1, 6, 3, 2, 5}
	fmt.Println(verifyPostorder(nums))
}
