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
//给定一个二叉搜索树（Binary Search Tree）
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	return add(root, &sum)
}

func add(root *TreeNode, sum *int) *TreeNode {
	if root == nil {
		return nil
	}
	add(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	add(root.Left, sum)
	return root
}

/**
	  5
	 / \
	2   13
 */
func main() {
	// -----------二叉树-----------------
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 13}
	root := &TreeNode{Val: 5, Left: node1, Right: node2}

	inOrderTraverse(root)
	fmt.Printf("\n")
	inOrderTraverse(convertBST(root))
}

func inOrderTraverse(tree *TreeNode) {
	if tree == nil {
		return
	}
	inOrderTraverse(tree.Left)
	fmt.Print(tree.Val, " ")
	inOrderTraverse(tree.Right)
}
