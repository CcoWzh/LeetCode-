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
func getKthFromEnd(head *ListNode, k int) *ListNode {
	if k < 1 { //如果k小于0，则直接返回nil
		return nil
	}
	fast, slow := head, head

	count := 0
	for count < k {
		fast = fast.Next
		if fast == nil { //如果k大于链表的长度，则直接返回全链表
			return head
		}
		count++
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}

func main() {
	//链表: 1->2->3->4->5->6

	link6 := ListNode{6, nil}
	link5 := ListNode{5, &link6}
	link4 := ListNode{4, &link5}
	link3 := ListNode{3, &link4}
	link2 := ListNode{2, &link3}
	link1 := ListNode{1, &link2}

	link1.PrintLink()
	getKthFromEnd(&link1, 7).PrintLink()
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
