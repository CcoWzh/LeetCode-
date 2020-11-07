package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * 83. 删除排序链表中的重复元素
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	result := head

	for head != nil {
		for head.Next != nil && head.Val == head.Next.Val {
			head.Next = head.Next.Next
		}
		head = head.Next
	}

	return result
}

func main() {
	//链表: 1->1->2->2->3->4

	link6 := ListNode{2, nil}
	link5 := ListNode{2, &link6}
	link4 := ListNode{2, &link5}
	link3 := ListNode{1, &link4}
	link2 := ListNode{1, &link3}
	link1 := ListNode{1, &link2}

	link1.PrintLink()
	newLink := deleteDuplicates1(&link1)
	newLink.PrintLink()

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

//更优解法
//func deleteDuplicates(head *ListNode) *ListNode {
//	if head == nil {
//		return head
//	}
//	p := head
//	preVal := p.Val
//	for p.Next != nil {
//		if p.Next.Val == preVal {
//			//最关键的，就是这个了，注意：
//			//不是 p = p.Next
//			p.Next = p.Next.Next
//		} else {
//			p = p.Next
//			preVal = p.Val
//		}
//	}
//	return head
//}

func deleteDuplicates1(head *ListNode) *ListNode {
	result := &ListNode{}

	res := result
	result.Next = head

	curVal := head.Val
	for result.Next != nil {
		if result.Next.Val == curVal {
			result.Next = result.Next.Next
		} else {
			result = result.Next
			curVal = result.Val
		}

	}

	return res.Next
}
