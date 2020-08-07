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
func sumEvenGrandparent(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Val%2 == 0 {
		return tow(root) + sumEvenGrandparent(root.Right) + sumEvenGrandparent(root.Left)
	} else {
		return sumEvenGrandparent(root.Right) + sumEvenGrandparent(root.Left)
	}
}

func tow(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for i := 0; i < 2 && len(queue) != 0; i++ {
		size := len(queue)
		for size > 0 {
			r := queue[0]
			queue = queue[1:]
			if r.Left != nil {
				queue = append(queue, r.Left)
			}
			if r.Right != nil {
				queue = append(queue, r.Right)
			}
			size--
		}
	}
	sum := 0
	for i := 0; i < len(queue); i++ {
		sum += queue[i].Val
	}

	return sum
}

func main() {
	// -----------二叉树-----------------
	node9 := &TreeNode{Val: 1}
	node8 := &TreeNode{Val: 9}
	node7 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 3, Right: node7}
	node3 := &TreeNode{Val: 2, Left: node8}
	node4 := &TreeNode{Val: 7, Left: node9}
	node5 := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 7, Left: node3, Right: node4}
	node2 := &TreeNode{Val: 8, Left: node5, Right: node6}
	root := TreeNode{Val: 6, Left: node1, Right: node2}
	fmt.Println(sumEvenGrandparent(&root))
}
