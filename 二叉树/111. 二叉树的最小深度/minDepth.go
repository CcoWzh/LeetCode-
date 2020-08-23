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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//广度优先搜索
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	//size,当前队列大小，count,当前层数
	size := 0
	count := 0
	for len(queue) != 0 {
		size = len(queue)
		count++
		for i := 0; i < size; i++ {
			r := queue[0]
			if r.Left != nil {
				queue = append(queue, r.Left)
			}
			if r.Right != nil {
				queue = append(queue, r.Right)
			}
			if r.Left == nil && r.Right == nil { //判断是否是叶子节点
				return count
			}
			queue = queue[1:] //出队列
		}
	}
	return count
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
	fmt.Println(minDepth(root))
}
