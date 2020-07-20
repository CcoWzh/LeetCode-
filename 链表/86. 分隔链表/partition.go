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
func partition(head *ListNode, x int) *ListNode {
	smallList, bigList := &ListNode{}, &ListNode{}
	smalll, big := smallList, bigList
	for head != nil {
		if head.Val < x {
			smallList.Next = &ListNode{Val: head.Val}
			smallList = smallList.Next
		} else {
			bigList.Next = &ListNode{Val: head.Val}
			bigList = bigList.Next
		}
		head = head.Next
	}
	//smallList,bigList现在都指向了链表的尾部
	smallList.Next = big.Next

	return smalll.Next
}

//1->4->3->2->5->2
func main() {
	link6 := ListNode{2, nil}
	link5 := ListNode{5, &link6}
	link4 := ListNode{2, &link5}
	link3 := ListNode{3, &link4}
	link2 := ListNode{4, &link3}
	link1 := ListNode{1, &link2}
	link1.PrintLink()

	partition(&link1, 3).PrintLink()
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
