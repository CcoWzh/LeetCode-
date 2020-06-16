package main

import "fmt"

//链表
type ListNode struct {
	Val  int
	Next *ListNode
}

//数组转单向链表
func arrayToList(nums []int) *ListNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	list := &ListNode{Val: nums[0]}
	temp := list
	for i := 1; i < n; i++ {
		temp.Next = &ListNode{Val: nums[i]}
		temp = temp.Next
	}

	return list
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

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	list := arrayToList(nums)
	list.PrintLink()
}
