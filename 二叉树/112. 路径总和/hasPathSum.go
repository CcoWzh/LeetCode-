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
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	stack := make([]*TreeNode, 0)
	path := make([]int, 0)
	cur := root
	stack = append(stack, cur)
	path = append(path, cur.Val)

	for len(stack) != 0 {
		cur, mySum := stack[0], path[0]
		stack, path = stack[1:], path[1:]
		if cur.Right == nil && cur.Left == nil && mySum == sum {
			return true
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
			path = append(path, mySum+cur.Left.Val)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
			path = append(path, mySum+cur.Right.Val)
		}
	}
	return false
}

//func hasPathSum(root *TreeNode, sum int) bool {
//	if root == nil {
//		return false
//	}
//	path := make([]int, 0)
//	curSum := 0
//	code := false
//	find(root, sum, curSum, path, code)
//	return code
//}
//
//func find(root *TreeNode, sum, curSum int, path []int, code bool) {
//	curSum += root.Val
//	path = append(path, root.Val)
//	if curSum == sum && root.Right == nil && root.Left == nil {
//		fmt.Println(path)
//		code = true
//	}
//
//	if root.Left != nil {
//		find(root.Left, sum, curSum, path, code)
//	}
//	if root.Right != nil {
//		find(root.Right, sum, curSum, path, code)
//	}
//	path = path[:len(path)-1]
//}

/**
	  1
	 / \
	4   8
   /   / \
  11  13  4
  /\       \
 7  2       1
 */
func main() {
	// -----------二叉树-----------------
	node8 := &TreeNode{Val: 7}
	node6 := &TreeNode{Val: 4, Left: nil, Right: node8}
	node7 := &TreeNode{Val: 2}
	node4 := &TreeNode{Val: 7}
	node5 := &TreeNode{Val: 13}
	node3 := &TreeNode{Val: 11, Left: node4, Right: node7}
	node1 := &TreeNode{Val: 4, Left: node3, Right: nil}
	node2 := &TreeNode{Val: 8, Left: node5, Right: node6}

	root := &TreeNode{Val: 1, Left: node1, Right: node2}
	sum := 22
	fmt.Println(hasPathSum(root, sum))
}
