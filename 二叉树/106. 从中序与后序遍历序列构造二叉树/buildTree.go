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
func buildTree(inorder []int, postorder []int) *TreeNode {
	n := len(inorder)
	if n == 0 {
		return nil
	} else if n == 1 {
		return &TreeNode{Val: postorder[0]}
	}
	rooVal := postorder[n-1]
	root := &TreeNode{Val: rooVal}
	index := 0
	for index < n {
		if inorder[index] == rooVal {
			break
		}
		index++
	}

	root.Left = buildTree(inorder[:index], postorder[:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:n-1])

	return root
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	root := buildTree(inorder, postorder)
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
