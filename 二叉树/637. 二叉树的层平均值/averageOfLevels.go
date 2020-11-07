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
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	res := make([]float64, 0)
	for len(queue) != 0 {
		size := len(queue)
		var sum float64
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			sum += float64(cur.Val)
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		res = append(res, sum/float64(size))
	}
	return res
}

/**
    3
   / \
  9  20
 / \   \
 2  1   7
    /
    15
 */
func main() {
	// -----------二叉树-----------------
	node3 := &TreeNode{Val: 15}
	node4 := &TreeNode{Val: 7}

	node5 := &TreeNode{Val: 2}
	node6 := &TreeNode{Val: 1, Left: node3}

	node1 := &TreeNode{Val: 20, Right: node4}
	node2 := &TreeNode{Val: 9, Left: node5, Right: node6}

	root := &TreeNode{Val: 3, Left: node1, Right: node2}
	fmt.Println(averageOfLevels(root))
}
