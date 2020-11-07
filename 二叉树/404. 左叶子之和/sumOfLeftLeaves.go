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
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	mark := make([]int, 0) //标志位，标记是否是左节点
	queue = append(queue, root)
	mark = append(mark, 0)

	res := 0
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur, sign := queue[0], mark[0]
			queue = queue[1:]
			mark = mark[1:]
			//首先判断是不是叶子节点，再判断是不是左节点
			if cur.Left == nil && cur.Right == nil && sign == 1 {
				res += cur.Val
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
				mark = append(mark, 1)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
				mark = append(mark, 0)
			}
		}
	}

	return res
}

/**
	  3
	 / \
	9  20
   / \ / \
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
	fmt.Println(sumOfLeftLeaves(&root))
}
