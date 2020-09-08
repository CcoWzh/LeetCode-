/**
使用栈操作
 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//给定一个二叉树，返回它的前序遍历。
func preorderTraversal(root *TreeNode) []int {
	var result []int
	stack := make([]*TreeNode, 0)
	cur := root
	stack = append(stack, cur)
	//原则:right先入栈，left后入栈（先打印left，后打印right）
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		result = append(result, node.Val)
		stack = stack[:len(stack)-1] //出栈

		if node.Right != nil { //right先入栈
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return result
}

/**
	  3
	 / \
	9  20
   / \ / \
  15 7 1  6
 */
func main() {
	// -----------二叉树-----------------
	node3 := &TreeNode{Val: 15}
	node4 := &TreeNode{Val: 7}
	node5 := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 9, Left: node3, Right: node4}
	node2 := &TreeNode{Val: 20, Left: node5}
	root := TreeNode{Val: 3, Left: node1, Right: node2}
	fmt.Println(preorderTraversal(&root))
}
