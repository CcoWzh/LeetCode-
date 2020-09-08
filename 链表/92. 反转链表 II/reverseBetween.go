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
// 1->2->3->4->5->6
//res i     j
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	//设置一个哨兵节点，防止m=1的情况
	//这是一个坑，不加的话，通不过所以的测试用例的
	result := &ListNode{}
	result.Next = head
	preHead := result

	for i := 1; i < m; i++ {
		preHead = head   //m之前的链指针
		head = head.Next //指向m的链指针
	}
	//需要转换的指针
	midHead := &ListNode{head.Val, head.Next}
	//开始反转链表
	temp_1 := midHead
	head = head.Next
	i := m
	for head != nil {
		if i >= n {
			break
		}
		temp := head.Next
		head.Next = midHead
		midHead = head
		head = temp
		i++
	}
	//连接链表
	preHead.Next = midHead
	//将反转后的链表，链接到结尾
	temp_1.Next = head

	return result.Next
}

func main() {
	//链表: 1->1->2->2->3->4

	link6 := ListNode{6, nil}
	link5 := ListNode{5, &link6}
	link4 := ListNode{4, &link5}
	link3 := ListNode{3, &link4}
	link2 := ListNode{2, &link3}
	link1 := ListNode{1, &link2}

	link1.PrintLink()
	newLink := reverseBetween(&link1, 1, 1)
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
