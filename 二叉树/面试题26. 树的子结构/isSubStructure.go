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
//约定空树不是任意一个树的子结构
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	//先要判断A，B是否为空
	if A == nil || B == nil {
		return false
	}

	result := false
	if A.Val == B.Val {
		result = isMatch(A, B)
	}
	if result {
		return true
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isMatch(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}

	if A.Val == B.Val {
		return isMatch(A.Left, B.Left) && isMatch(A.Right, B.Right)
	} else {
		return false
	}
}

/**
	  3
	 / \
	9  20
   / \ / \
  15 7 1  6
 */
func main() {
	// -----------二叉树-----------------
	node6 := &TreeNode{Val: 6}
	node3 := &TreeNode{Val: 15}
	node4 := &TreeNode{Val: 7}
	node5 := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 9, Left: node3, Right: node4}
	node2 := &TreeNode{Val: 20, Left: node5, Right: node6}
	root := &TreeNode{Val: 3, Left: node1, Right: node2}
	fmt.Println(isSubStructure(root, node2))
}
