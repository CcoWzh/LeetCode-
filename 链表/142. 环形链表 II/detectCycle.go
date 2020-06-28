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
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	//判断是否有环
	for fast != nil {
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			break
		}
	}
	//这里要附加一个判断，判断是否有环，还是直接就是nil跳出的for循环
	if slow != fast {
		return nil
	}
	//判断环在哪里
	slow = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}

func main() {
	//链表: 1->2->3->4->5->2
	link5 := ListNode{5, nil}
	link4 := ListNode{4, &link5}
	link3 := ListNode{3, &link4}
	link2 := ListNode{2, &link3}
	link1 := ListNode{1, &link2}
	link5.Next = &link2

	newLink := detectCycle(&link1)
	fmt.Println(newLink.Val)
}
