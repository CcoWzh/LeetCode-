package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * 82. 删除排序链表中的重复元素 II
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	preHead := &ListNode{}
	result := preHead

	for head != nil {
		if head.Next == nil {
			preHead.Next = &ListNode{head.Val, nil}
			preHead = preHead.Next
			break
		}

		if head.Val != head.Next.Val {
			preHead.Next = &ListNode{head.Val, nil}
			preHead = preHead.Next
			head = head.Next
		} else {
			curVal := head.Val
			for head.Val == curVal && head.Next != nil {
				head = head.Next
			}

			if head.Next == nil && head.Val == curVal {
				break
			}
		}
	}

	return result.Next
}

func main() {
	//链表: 1->2->3->3->4->4->5

	link7 := ListNode{4, nil}
	link6 := ListNode{4, &link7}
	link5 := ListNode{4, &link6}
	link4 := ListNode{3, &link5}
	link3 := ListNode{1, &link4}
	link2 := ListNode{1, &link3}
	link1 := ListNode{1, &link2}

	link1.PrintLink()
	newLink := deleteDuplicates(&link1)
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
