/**
以前可以使用pre = head，原因是，链表结构没有改变
但是，一旦反转了链表或者破坏了整个链表的结构，则这种赋值、拷贝都会发生改变
这点，要搞清楚，切记切记！！！
 */
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
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	//1->2->2->1
	tail := reverse(slow)
	//tail ==>2->1
	//head.PrintLink()  ==> 1->2->2
	//slow此时，本来就是指向nil
	//slow.Next = nil  //这个，确实是不必要的操作
	for head != nil && tail != nil {
		if head.Val != tail.Val {
			return false
		}
		head = head.Next
		tail = tail.Next
	}

	return true
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	for head != nil {
		temp := head.Next
		head.Next = pre
		pre = head
		head = temp
	}
	return pre
}

func main() {
	//链表: 1->2->3->2->1
	link5 := ListNode{1, nil}
	link4 := ListNode{2, &link5}
	//link3 := ListNode{3, &link4}
	link2 := ListNode{2, &link4}
	link1 := ListNode{1, &link2}

	link1.PrintLink()

	fmt.Println(isPalindrome(&link1))
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
