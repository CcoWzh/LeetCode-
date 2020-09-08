package main

import "fmt"

//链表
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	mark := 0
	res := &ListNode{}
	head := res

	for l1 != nil || l2 != nil {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}
		sum := a + b + mark
		mark = sum / 10
		sum = sum % 10

		res.Next = &ListNode{Val: sum}
		res = res.Next
	}
	if mark == 1 {
		res.Next = &ListNode{Val: 1}
		res = res.Next
	}

	return head.Next
}

//(2 -> 4 -> 3) + (5 -> 6 -> 4)
//7 -> 0 -> 8
func main() {
	link5 := ListNode{9, nil}
	link4 := ListNode{9, &link5}
	l1 := ListNode{9, &link4}

	link2 := ListNode{9, nil}
	link1 := ListNode{9, &link2}
	l2 := ListNode{9, &link1}

	l1.PrintLink()
	l2.PrintLink()
	l3 := addTwoNumbers(&l1, &l2)
	l3.PrintLink()
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
