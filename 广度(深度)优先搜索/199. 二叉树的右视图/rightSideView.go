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
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:] //出队列
			if i == size-1 {
				res = append(res, cur.Val)
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}

	return res
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
	node3 := &TreeNode{Val: 15}
	node4 := &TreeNode{Val: 7}
	node5 := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 9, Left: node3, Right: node4}
	node2 := &TreeNode{Val: 20, Left: node5}
	root := TreeNode{Val: 3, Left: node1, Right: node2}
	fmt.Println(rightSideView(&root))
}
