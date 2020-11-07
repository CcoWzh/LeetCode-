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
func swapPairs(head *ListNode) *ListNode {
	res := &ListNode{}
	other := res
	for head != nil {
		if head.Next == nil {
			other.Next = head
			break
		}
		temp := head.Next.Next //缓存
		//获得首、尾的链表指针
		first, second := reverse(head)
		//开始反转
		other.Next = second
		other = first //这一步，就是直接将“结果指针”赋值到下一个反转的位置
		//first.Next = temp //这一步似乎是不必要的
		head = temp
	}

	return res.Next
}

//两个元素反转,返回两个指针：
//反转后的尾部first，反转后的首部second
func reverse(head *ListNode) (first *ListNode, second *ListNode) {
	first = head
	second = &ListNode{}
	for i := 0; i < 2; i++ {
		next := head.Next
		head.Next = second
		second = head
		head = next
	}
	return first, second
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

	swapPairs(&link1).PrintLink()
	//s, e := reverse(&link1)
	//e.PrintLink()
	//s.PrintLink()
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
