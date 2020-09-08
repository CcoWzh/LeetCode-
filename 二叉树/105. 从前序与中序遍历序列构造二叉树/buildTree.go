package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	n := len(preorder)
	if n == 0 {
		return nil
	} else if n == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	base := preorder[0]
	index := 0
	for index < len(inorder) {
		if inorder[index] == base {
			break
		}
		index++
	}
	root := &TreeNode{Val: base}
	root.Left = buildTree(preorder[1:index+1], inorder[:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])
	return root
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	fmt.Println(inorderTraversal(root))
}

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
		//fmt.Println(cur.Val)
		stack = stack[:len(stack)-1] // 出栈
		result = append(result, cur.Val)
		cur = cur.Right
	}
	return result
}
