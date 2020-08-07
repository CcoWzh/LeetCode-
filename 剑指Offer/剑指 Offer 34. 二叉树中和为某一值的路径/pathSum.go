package main

import (
	"fmt"
)

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
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 0)
	dfs(root, []int{}, 0, sum, &result)
	return result
}

func dfs(node *TreeNode, path []int, cur, sum int, result *[][]int) {
	if node == nil {
		return
	}
	cur += node.Val
	path = append(path, node.Val)
	if node.Left == nil && node.Right == nil && cur == sum {
		*result = append(*result, path)
	}

	temp := make([]int, len(path))
	copy(temp,path)
	dfs(node.Left, temp, cur, sum, result)
	dfs(node.Right, temp, cur, sum, result)
	path = path[:len(path)-1]
	cur -= node.Val
}

/**
             5
           /  \
          4    8
         / \  / \
        5  7  1  2
 */
func main() {
	// -----------二叉树-----------------
	node3 := &TreeNode{Val: 5}
	node4 := &TreeNode{Val: 7}
	node5 := &TreeNode{Val: 1}
	node6 := &TreeNode{Val: 2}
	node1 := &TreeNode{Val: 4, Left: node3, Right: node4}
	node2 := &TreeNode{Val: 8, Left: node5, Right: node6}
	root := TreeNode{Val: 5, Left: node1, Right: node2}
	fmt.Println(pathSum(&root, 14))
}
