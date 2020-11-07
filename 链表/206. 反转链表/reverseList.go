package main

import "fmt"

//链表
type ListNode struct {
	Val  int
	Next *ListNode
}

//反转链表
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prehead *ListNode
	for head != nil {
		next := head.Next
		head.Next = prehead
		prehead = head
		head = next
	}

	return prehead
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

//链表: 1->2->3->4->5
func main() {
	link5 := ListNode{5, nil}
	link4 := ListNode{4, &link5}
	link3 := ListNode{3, &link4}
	link2 := ListNode{2, &link3}
	link1 := ListNode{1, &link2}

	link1.PrintLink()
	//s复制的是一个节点
	s := link2
	s.Val = 100
	s.PrintLink()
	link1.PrintLink()
	fmt.Println("反转链表：")
	reverseList(&link1).PrintLink()
	fmt.Println("s的值：")
	s.PrintLink()
}
