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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) != 0 {
		size := len(queue)
		temp := make([]int, size)
		for i := 0; i < size; i++ {
			r := queue[0]
			queue = queue[1:]
			temp[i] = r.Val
			if r.Left != nil {
				queue = append(queue, r.Left)
			}
			if r.Right != nil {
				queue = append(queue, r.Right)
			}
		}
		res = append(res, temp)
	}

	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return res
}

/**
    3
   / \
  20  9
 /   / \
 7  2   1
	     \
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
	fmt.Println(levelOrderBottom(root))
}
