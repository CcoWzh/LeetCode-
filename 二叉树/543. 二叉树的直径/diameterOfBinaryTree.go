package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	res = 0
	def(root)
	return res
}

//返回以该节点为根的子树的深度
func def(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := def(root.Left)
	r := def(root.Right)
	res = max(res, l+r)

	return max(l, r) + 1
}

func max(a, b int) int {
	if a < b {
		a = b
	}
	return a
}

/**
	  3
	 / \
	9  20
   / \ /
  15 7 1
 */
func main() {
	// -----------二叉树-----------------
	//node3 := &TreeNode{Val: 15}
	//node4 := &TreeNode{Val: 7}
	//node5 := &TreeNode{Val: 1}
	//node1 := &TreeNode{Val: 9, Left: node3, Right: node4}
	//node2 := &TreeNode{Val: 20, Left: node5}
	//root := TreeNode{Val: 3, Left: node1, Right: node2}
	root := TreeNode{}
	fmt.Println(diameterOfBinaryTree(&root))
}
