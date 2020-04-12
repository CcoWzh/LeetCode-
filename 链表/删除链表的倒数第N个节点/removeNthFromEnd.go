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
//哨兵节点+双指针法
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	result := &ListNode{}
	result.Next = head
	var pre *ListNode
	cur := result
	//不是head.Next != nil
	for i := 1; head != nil; i++ {
		if i >= n {
			pre = cur
			cur = cur.Next
		}
		head = head.Next
	}
	pre.Next = pre.Next.Next

	return result.Next
}

func (head *ListNode) GetLength() int {
	iterator := head

	length := 0
	fmt.Print("LinkNode is :")
	for iterator.Next != nil {
		fmt.Print(iterator.Val)
		fmt.Print("->")
		length ++
		iterator = iterator.Next
	}
	fmt.Print(iterator.Val)
	fmt.Println()
	return length + 1
}

//链表: 1->2->3->4->5
func main() {
	link5 := ListNode{5, nil}
	link4 := ListNode{4, &link5}
	link3 := ListNode{3, &link4}
	link2 := ListNode{2, &link3}
	link1 := ListNode{1, &link2}

	fmt.Println("link is ", link1.GetLength())
	removeNthFromEnd(&link1, 2)
	fmt.Println("link is ", link1.GetLength())
}
