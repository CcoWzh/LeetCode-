package main

import (
	"fmt"
	"strconv"
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
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	path := make([]string, 0) //保存当前遍历的路径
	queue = append(queue, root)
	path = append(path, strconv.Itoa(root.Val))

	res := make([]string, 0)
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			r := queue[0]
			queue = queue[1:]
			curPath := path[0]
			path = path[1:] //和节点一起，出队列
			//遍历到叶子节点，则输出结果
			if r.Left == nil && r.Right == nil {
				res = append(res, curPath)
			}
			//如果是非叶子的话，则入队列
			if r.Left != nil {
				queue = append(queue, r.Left)
				path = append(path, curPath+"->"+strconv.Itoa(r.Left.Val))
			}
			if r.Right != nil {
				queue = append(queue, r.Right)
				path = append(path, curPath+"->"+strconv.Itoa(r.Right.Val))
			}
		}

	}

	return res
}

/**
		3
	   9  20
      15 7  1
 */
func main() {
	// -----------二叉树-----------------
	node3 := &TreeNode{Val: 15}
	node4 := &TreeNode{Val: 7}
	node5 := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 9, Left: node3, Right: node4}
	node2 := &TreeNode{Val: 20, Left: node5}
	root := TreeNode{Val: 3, Left: node1, Right: node2}
	fmt.Println(binaryTreePaths(&root))
}
