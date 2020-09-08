package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//归并排序，其实，似乎和“合并两个有序链表”相似
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//找到中间的节点,快慢指针
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	//将当前链表一分为二
	right := slow.Next
	left := head
	slow.Next = nil
	//归并排序，递归调用
	l1 := sortList(left)
	l2 := sortList(right)
	//开始归并，即，合并两个有序链表
	preHead := &ListNode{}
	res := preHead
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			preHead.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		} else {
			preHead.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		}
		preHead = preHead.Next
	}
	if l1 != nil {
		preHead.Next = l1
	} else if l2 != nil {
		preHead.Next = l2
	}

	return res.Next
}

func main() {
	//链表: -1->5->3->4->0->2

	link6 := ListNode{2, nil}
	link5 := ListNode{0, &link6}
	link4 := ListNode{4, &link5}
	link3 := ListNode{3, &link4}
	link2 := ListNode{5, &link3}
	link1 := ListNode{-1, &link2}

	link1.PrintLink()
	newLink := sortList(&link1)
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
