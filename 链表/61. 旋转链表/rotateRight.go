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
func rotateRight(head *ListNode, k int) *ListNode {
	lenght := 0
	fast := head
	for fast != nil { //得到链表的长度
		lenght++
		fast = fast.Next
	}
	k = k % lenght
	if k == 0 {
		return head
	}
	//快慢指针，快指针先多走k步
	fast = head
	slow := head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	//开始旋转
	res := slow.Next
	slow.Next = nil
	fast = res
	for fast.Next != nil {
		fast = fast.Next
	}
	fast.Next = head

	return res
}

func main() {
	//1->2->3->4->5
	link1 := ListNode{5, nil}
	link2 := ListNode{4, &link1}
	link5 := ListNode{3, &link2}
	link4 := ListNode{2, &link5}
	l1 := ListNode{1, &link4}

	l1.PrintLink()
	l3 := rotateRight(&l1, 9)
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
