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
func reverseKGroup(head *ListNode, k int) *ListNode {
	count := 0
	end := head
	sentinel := &ListNode{}
	result := sentinel
	for end != nil {
		start := end
		for count < k {
			if end == nil {
				sentinel.Next = start
				return result.Next
			}
			end = end.Next
			count++
		}
		s, e := reverse(start, k)
		sentinel.Next = e
		sentinel = s
		count = 0
	}

	return result.Next
}

func reverse(start *ListNode, k int) (s *ListNode, e *ListNode) {
	s, count := start, 0
	for count < k && start != nil{
		temp := start.Next
		start.Next = e
		e = start
		start = temp
		count++
	}
	return s, e
}

//1->2->3->4->5->6
func main() {
	//link6 := ListNode{6, nil}
	link5 := ListNode{5, nil}
	link4 := ListNode{4, &link5}
	link3 := ListNode{3, &link4}
	link2 := ListNode{2, &link3}
	link1 := ListNode{1, &link2}
	link1.PrintLink()

	reverseKGroup(&link1, 5).PrintLink()
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
