package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	res := 0
	for head != nil {
		res = res*2 + head.Val
		head = head.Next
	}
	return res
}

func main() {
	link3 := ListNode{1, nil}
	link2 := ListNode{0, &link3}
	link1 := ListNode{1, &link2}

	fmt.Println(getDecimalValue(&link1))
}
