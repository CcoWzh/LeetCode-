package main

import (
	"fmt"
	"math"
)

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
func isValidBST(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)
	base := math.MinInt64
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		top := stack[len(stack)-1]   //栈顶元素
		stack = stack[:len(stack)-1] //出栈
		if top.Val <= base { //即，要保证中序遍历结果是递增的
			return false
		}
		base = top.Val
		root = top.Right
	}

	return true
}

/**
    5
   / \
  1   7
     / \
    3   8
 */
func main() {
	// -----------二叉树-----------------
	node4 := &TreeNode{Val: 8}
	node5 := &TreeNode{Val: 3}
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 7, Left: node5, Right: node4}
	root := TreeNode{Val: 5, Left: node1, Right: node2}
	fmt.Println(isValidBST(&root))
}
