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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l := layers(root.Left)
	r := layers(root.Right)
	fmt.Println(l, r)
	//确保---任意节点
	return abs(l-r) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

//求绝对值
func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

//计算二叉树的层数
func layers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := layers(root.Left)
	r := layers(root.Right)

	return max(l, r) + 1
}

//求最大值
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
	fmt.Println(isBalanced(&root))
}
