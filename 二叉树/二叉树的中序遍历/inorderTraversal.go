package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//给定一个二叉树，返回它的中序 遍历。
//使用切片作为栈 stack = stack[:l-1]
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

func main() {
	// -----------二叉树-----------------
	node3 := &TreeNode{Val: 15}
	node4 := &TreeNode{Val: 7}
	node5 := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 9, Left: node3, Right: node4}
	node2 := &TreeNode{Val: 20, Left: node5}
	root := TreeNode{Val: 3, Left: node1, Right: node2}
	fmt.Println(inorderTraversal(&root))
}
