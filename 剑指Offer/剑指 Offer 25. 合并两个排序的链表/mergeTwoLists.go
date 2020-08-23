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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	temp := &ListNode{}
	result := temp

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			temp.Next = &ListNode{Val: l1.Val}
			temp = temp.Next
			l1 = l1.Next
		} else {
			temp.Next = &ListNode{Val: l2.Val}
			temp = temp.Next
			l2 = l2.Next
		}
	}

	if l1 != nil {
		temp.Next = l1
	} else if l2 != nil {
		temp.Next = l2
	}

	return result.Next
}

func main() {
	//链表: 1->4->8
	//     2->5->6

	link6 := ListNode{6, nil}
	link5 := ListNode{5, &link6}
	link4 := ListNode{2, &link5}
	link3 := ListNode{8, nil}
	link2 := ListNode{4, &link3}
	link1 := ListNode{1, &link2}

	mergeTwoLists(&link1, &link4).PrintLink()
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
