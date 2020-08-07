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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a < b {
		a = b
	}
	return a
}

/**
             3
           /  \
          9   20
         /      \
        15       1
        /         \
       2           7
 */
func main() {
	// -----------二叉树-----------------
	node4 := &TreeNode{Val: 7}
	node5 := &TreeNode{Val: 1, Right: node4}
	node6 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 15, Left: node6}
	node1 := &TreeNode{Val: 9, Left: node3}
	node2 := &TreeNode{Val: 20, Right: node5}
	root := TreeNode{Val: 3, Left: node1, Right: node2}
	fmt.Println(maxDepth(&root))
}
