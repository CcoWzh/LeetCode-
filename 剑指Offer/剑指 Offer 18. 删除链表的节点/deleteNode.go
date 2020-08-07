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
func deleteNode(head *ListNode, val int) *ListNode {
	per := &ListNode{}
	result := per

	per.Next = head
	//per从空节点开始，为了防止删除的是头节点
	for per.Next != nil {
		if per.Next.Val == val {
			per.Next = per.Next.Next
			break
		}
		per = per.Next
	}

	return result.Next
}

func main() {
	//链表: 1->2->3->4->5->5

	link6 := ListNode{6, nil}
	link5 := ListNode{5, &link6}
	link4 := ListNode{4, &link5}
	link3 := ListNode{3, &link4}
	link2 := ListNode{2, &link3}
	link1 := ListNode{1, &link2}

	link1.PrintLink()
	deleteNode(&link1, 5).PrintLink()
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
