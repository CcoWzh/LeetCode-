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
	indexA, indexB := 0, 0
	l1, l2 := headA, headB
	for l1 != nil {
		indexA++
		l1 = l1.Next
	}
	for l2 != nil {
		indexB++
		l2 = l2.Next
	}

	if indexA > indexB {
		for i := 0; i < indexA-indexB; i++ {
			headA = headA.Next
		}
	} else {
		for i := 0; i < indexB-indexA; i++ {
			headB = headB.Next
		}
	}

	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}
		headA = headA.Next
		headB = headB.Next
	}
	return nil
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
