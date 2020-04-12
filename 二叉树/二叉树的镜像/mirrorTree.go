package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//面试题27. 二叉树的镜像
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	//交换左右节点
	root.Right, root.Left = root.Left, root.Right
	mirrorTree(root.Left)
	mirrorTree(root.Right)
	return root
}

//前序遍历：根 -----> 左子树 -----> 右子树
//递归实现
func preOrderTraverse(tree *TreeNode) {
	if tree == nil {
		return
	} else {
		fmt.Print(tree.Val, " ")
		preOrderTraverse(tree.Left)
		preOrderTraverse(tree.Right)
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
	preOrderTraverse(root)
	fmt.Println("\n")
	preOrderTraverse(mirrorTree(root))
}
