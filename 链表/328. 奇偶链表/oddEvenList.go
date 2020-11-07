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
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	first, second := head, head.Next
	res, temp := first, second //保存原始指针

	for first != nil && second != nil {
		if second.Next != nil { //保证节点为偶数时，编程正确
			first.Next = second.Next
			first = first.Next
		} else {
			break
		}
		if first != nil { //保证节点为奇数时，编程正确
			second.Next = first.Next
			second = second.Next
		} else {
			break
		}
	}

	first.Next = temp
	return res
}

//1->2->3->4->5
func main() {
	//link6 := ListNode{6, nil}
	link5 := ListNode{5, nil}
	link4 := ListNode{4, &link5}
	link3 := ListNode{3, &link4}
	link2 := ListNode{2, &link3}
	link1 := ListNode{1, &link2}
	link1.PrintLink()

	oddEvenList(&link1).PrintLink()
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
