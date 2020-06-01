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
func isSymmetric(root *TreeNode) bool {
	return help(root,root)
}

func help(root1,root2 *TreeNode) bool {
	if root1 == nil && root2 == nil{
		return true
	}
	if root1 == nil || root2 == nil{
		return false
	}

	if root1.Val != root2.Val{
		return false
	}

	return help(root1.Left,root2.Right) && help(root1.Right,root2.Left)
}



/**
	  1
	 / \
	2   2
   / \ / \
  3  4 4  3
 */
func main() {
	// -----------二叉树-----------------
	node6 := &TreeNode{Val: 3}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 4}
	node1 := &TreeNode{Val: 2, Left: node3, Right: node4}
	node2 := &TreeNode{Val: 2, Left: node5, Right: node6}

	//node1 := &TreeNode{Val: 2, Left: nil, Right: node3}
	//node2 := &TreeNode{Val: 2, Left: nil, Right: node6}

	root := &TreeNode{Val: 1, Left: node1, Right: node2}
	fmt.Println(isSymmetric(root))
}
