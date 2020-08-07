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
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

/**
	  1
	 / \
	3   4
   / \ / \
       5  6
 */
func main() {
	// -----------二叉树-----------------
	node1 := &TreeNode{Val: 3}
	node2 := &TreeNode{Val: 5}
	node4 := &TreeNode{Val: 6}
	node3 := &TreeNode{Val: 4, Left: node2, Right: node4}
	p := &TreeNode{Val: 1, Left: node1, Right: node3}
	q := &TreeNode{Val: 1, Left: node1, Right: node3}

	fmt.Println(isSameTree(p, q))
}
