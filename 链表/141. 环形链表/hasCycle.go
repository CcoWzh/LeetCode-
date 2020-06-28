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
 * 或者，可以使用map，可以降低时间复杂度
 */
func hasCycle(head *ListNode) bool {
	fast, slow := head, head

	for fast != nil {
		if fast.Next == nil || fast.Next.Next == nil {
			break
		}
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			return true
		}
	}

	return false
}

func main() {
	//链表: 1->2->3->4->5->2

	link5 := ListNode{5, nil}
	link4 := ListNode{4, &link5}
	link3 := ListNode{3, &link4}
	link2 := ListNode{2, &link3}
	link1 := ListNode{1, &link2}
	link5.Next = &link2

	fmt.Println(hasCycle(&link1))
}
