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
func removeDuplicateNodes(head *ListNode) *ListNode {
	result := head
	m := make(map[int]bool)

	for head != nil {
		m[head.Val] = true
		if head.Next != nil && m[head.Next.Val] == true {
			if head.Next.Next == nil {
				head.Next = nil
				break
			}
			head.Next = head.Next.Next
			continue
		}
		head = head.Next
	}

	return result
}

func main() {
	//链表: 1->3->2->2->4->3

	link6 := ListNode{3, nil}
	link5 := ListNode{4, &link6}
	link4 := ListNode{2, &link5}
	link3 := ListNode{2, &link4}
	link2 := ListNode{3, &link3}
	link1 := ListNode{1, &link2}

	link1.PrintLink()
	newLink := removeDuplicateNodes(&link1)
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
