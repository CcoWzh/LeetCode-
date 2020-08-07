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
//递归求法
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//fmt.Println(root.Val)
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(m, n int) int {
	if m > n {
		return m
	} else {
		return n
	}
}

func main() {
	// -----------二叉树-----------------
	node3 := &TreeNode{Val: 15}
	node4 := &TreeNode{Val: 7}
	node5 := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 9, Left: node3, Right: node4}
	node2 := &TreeNode{Val: 20, Left: node5}
	root := TreeNode{Val: 3, Left: node1, Right: node2}
	fmt.Println(maxDepth(&root))
	fmt.Println(maxDepth1(&root))
}

//广度优先遍历方法
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	res := 0
	for len(queue) != 0 {
		size := len(queue)
		for size > 0 {
			//出队列
			r := queue[0]
			queue = queue[1:]
			//入队列
			if r.Left != nil {
				queue = append(queue, r.Left)
			}
			if r.Right != nil {
				queue = append(queue, r.Right)
			}
			size--
		}
		res++
	}

	return res
}
