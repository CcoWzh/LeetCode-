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

	prehead := head
	cur := head.Next
	prehead.Next = nil //步骤可不能调换了
	for cur != nil {
		next := cur.Next
		cur.Next = prehead
		prehead = cur
		cur = next
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
	reverseList(&link1).PrintLink()
}
