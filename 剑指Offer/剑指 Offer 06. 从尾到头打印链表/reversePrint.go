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
func reversePrint(head *ListNode) []int {
	stack := make([]int, 0)
	for head != nil {
		//入栈
		stack = append(stack, head.Val)
		head = head.Next
	}
	n := len(stack)
	result := make([]int, n)
	//其中，可以直接反转stack数组的
	for i := 0; i < n; i++ {
		l := len(stack)
		result[i] = stack[l-1]
		//出栈
		stack = stack[:l-1]
		//fmt.Println(stack)
	}

	return result
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
	fmt.Println(reversePrint(&link1))
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
