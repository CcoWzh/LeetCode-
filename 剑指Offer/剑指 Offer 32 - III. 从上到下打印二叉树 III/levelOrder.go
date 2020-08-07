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
//广度优先遍历
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	res := make([][]int, 0)
	flag := 1
	for len(queue) != 0 {
		temp := []*TreeNode{}
		tempRes := []int{}
		n := len(queue)
		for i := 0; i < n; i++ {
			r := queue[i]
			//其实，广度遍历的顺序不需要改变，只需要改变承接的值的顺序
			if flag%2 == 0 {
				s := queue[n-i-1]
				tempRes = append(tempRes, s.Val)
			} else {
				tempRes = append(tempRes, r.Val)
			}
			if r.Left != nil {
				temp = append(temp, r.Left)
			}
			if r.Right != nil {
				temp = append(temp, r.Right)
			}
		}
		flag++
		queue = temp
		res = append(res, tempRes)
	}
	//fmt.Println(flag)
	return res
}

/**
             3
           /  \
          9   20
         / \  / \
        15 7  1  2
 */
func main() {
	// -----------二叉树-----------------
	node3 := &TreeNode{Val: 15}
	node4 := &TreeNode{Val: 7}
	node5 := &TreeNode{Val: 1}
	node6 := &TreeNode{Val: 2}
	node1 := &TreeNode{Val: 9, Left: node3, Right: node4}
	node2 := &TreeNode{Val: 20, Left: node5, Right: node6}
	root := TreeNode{Val: 3, Left: node1, Right: node2}
	fmt.Println(levelOrder(&root))
}
