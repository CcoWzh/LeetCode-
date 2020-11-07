package main

import "fmt"

//876. 链表的中间结点
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
//这题可以使用快慢指针来解决
//用两个指针 slow 与 fast 一起遍历链表。
//slow 一次走一步，fast 一次走两步。那么当 fast 到达链表的末尾时，slow 必然位于中间。
func middleNode(head *ListNode) *ListNode {
	solw, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		solw = solw.Next
	}

	return solw
}

//链表: 1->2->3->4->5
func main() {
	link6 := ListNode{6, nil}
	link5 := ListNode{5, &link6}
	link4 := ListNode{4, &link5}
	link3 := ListNode{3, &link4}
	link2 := ListNode{2, &link3}
	link1 := ListNode{1, &link2}

	link1.PrintLink()
	middle := middleNode(&link1)
	middle.PrintLink()
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
