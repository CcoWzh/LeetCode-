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
//1 ≤ k ≤ 二叉搜索树元素个数 ,似乎默认了不会是空节点
func kthLargest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	mid := make([]int, 0)
	middleOrder(root, &mid)
	return mid[len(mid)-k]
}

func middleOrder(root *TreeNode, mid *[]int) {
	if root == nil {
		return
	} else {
		middleOrder(root.Left, mid)
		*mid = append(*mid, root.Val)
		middleOrder(root.Right, mid)
	}
}

/**
	  5
	 / \
	3   7
   / \ / \
  2  4 6  8
 */
func main() {
	// -----------二叉树-----------------
	node6 := &TreeNode{Val: 8}
	node3 := &TreeNode{Val: 2}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 6}
	node1 := &TreeNode{Val: 3, Left: node3, Right: node4}
	node2 := &TreeNode{Val: 7, Left: node5, Right: node6}
	root := &TreeNode{Val: 5, Left: node1, Right: node2}

	fmt.Println(kthLargest(root, 3))
}
