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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//栈
	stackA := make([]*ListNode, 0)
	stackB := make([]*ListNode, 0)
	c := headA

	for headA != nil {
		stackA = append(stackA, headA)
		headA = headA.Next
	}
	for headB != nil {
		stackB = append(stackB, headB)
		headB = headB.Next
	}

	for len(stackA) != 0 && len(stackB) != 0 {
		a := stackA[len(stackA)-1]
		b := stackB[len(stackB)-1]
		if a != b {
			break
		}
		//出栈
		stackA = stackA[:len(stackA)-1]
		stackB = stackB[:len(stackB)-1]
	}

	if len(stackA) != 0 {
		return stackA[len(stackA)-1].Next
	} else if len(stackB) != 0 {
		return stackB[len(stackB)-1].Next
	} else {
		return c
	}
}

/**
   4->1->8->4->5
5->0->1->8->4->5
 */
func main() {

	link6 := ListNode{5, nil}
	link5 := ListNode{4, &link6}
	link3 := ListNode{8, &link5}
	link4 := ListNode{1, &link3}
	l1 := ListNode{4, &link4}
	l1.PrintLink()

	link2 := ListNode{1, &link3}
	link1 := ListNode{0, &link2}
	l2 := ListNode{5, &link1}
	l2.PrintLink()

	getIntersectionNode(&l1, &l2).PrintLink()
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
