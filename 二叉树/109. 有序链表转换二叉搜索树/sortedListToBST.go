package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	tree := make([]int, 0)
	for head != nil {
		tree = append(tree, head.Val)
		head = head.Next
	}
	//构建平衡树
	return buildTree(tree, 0, len(tree)-1)
}

func buildTree(tree []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := (start + end) / 2

	root := &TreeNode{Val: tree[mid]}
	root.Left = buildTree(tree, start, mid-1)
	root.Right = buildTree(tree, mid+1, end)

	return root
}

func main() {
	//链表: 1->1->2->2->3->4
	link5 := ListNode{9, nil}
	link4 := ListNode{5, &link5}
	link3 := ListNode{0, &link4}
	link2 := ListNode{-3, &link3}
	link1 := ListNode{-10, &link2}
	root := sortedListToBST(&link1)
	fmt.Println(inorderTraversal(root))
	fmt.Println(root.Val,root.Left.Val,root.Right.Val)
}

//中序遍历
func inorderTraversal(root *TreeNode) []int {
	var result []int
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur) //入栈
			cur = cur.Left
		}
		cur = stack[len(stack)-1:][0]
		stack = stack[:len(stack)-1] // 出栈
		result = append(result, cur.Val)
		cur = cur.Right
	}
	return result
}
