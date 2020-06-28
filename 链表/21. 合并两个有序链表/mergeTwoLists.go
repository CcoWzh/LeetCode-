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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := &ListNode{}
	result := prehead

	for l1 != nil && l2 != nil {

		if l1.Val < l2.Val {
			prehead.Next = l1
			l1 = l1.Next
		} else {
			prehead.Next = l2
			l2 = l2.Next
		}
		prehead = prehead.Next //指向下一个指针
	}

	if l1 != nil {
		prehead.Next = l1
	}
	if l2 != nil {
		prehead.Next = l2
	}

	return result.Next
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

//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
func main() {
	link5 := ListNode{4, nil}
	link4 := ListNode{2, &link5}
	l1 := ListNode{1, &link4}

	link2 := ListNode{4, nil}
	link1 := ListNode{3, &link2}
	l2 := ListNode{1, &link1}

	l1.PrintLink()
	l2.PrintLink()
	l3 := mergeTwoLists(&l1, &l2)
	l3.PrintLink()
}
