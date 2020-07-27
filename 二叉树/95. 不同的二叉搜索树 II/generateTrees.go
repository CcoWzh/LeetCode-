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
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}

	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	var root []*TreeNode
	for i := start; i <= end; i++ {
		left := helper(start, i-1)
		right := helper(i+1, end)

		for _, x := range left {
			for _, y := range right {
				myRoot := &TreeNode{Val: i}
				myRoot.Left = x
				myRoot.Right = y
				root = append(root, myRoot)
			}
		}
	}
	return root
}

func main() {
	root := generateTrees(3)
	fmt.Println(len(root))
}
