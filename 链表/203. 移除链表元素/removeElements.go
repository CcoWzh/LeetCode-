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
func removeElements(head *ListNode, val int) *ListNode {
	pre := &ListNode{}
	pre.Next = head
	res := pre

	for pre != nil {
		for pre.Next != nil && pre.Next.Val == val {
			pre.Next = pre.Next.Next
		}
		pre = pre.Next
	}

	return res.Next
}

//链表: 1->2->3->4->5->6
func main() {
	link6 := ListNode{6, nil}
	link5 := ListNode{6, &link6}
	link4 := ListNode{6, &link5}
	link3 := ListNode{6, &link4}
	link2 := ListNode{2, &link3}
	link1 := ListNode{1, &link2}

	link1.PrintLink()
	middle := removeElements(&link1, 6)
	middle.PrintLink()
}

func (head *ListNode) PrintLink() {
	iterator := head
	length := 0
	fmt.Print("LinkNode is :")
	for iterator != nil {
		fmt.Print(iterator.Val)
		length ++
		iterator = iterator.Next
		if iterator != nil {
			fmt.Print("->")
		}
	}
	fmt.Println(" , length is ", length)
}
