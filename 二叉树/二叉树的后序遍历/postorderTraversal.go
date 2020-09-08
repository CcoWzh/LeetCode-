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

/**
前序遍历要求：root–>left–>right
后序遍历要求：left–>right–root
可以理解为先将先序遍历的left和right顺序互换，然后整体反转，得到后序遍历结果。
 */
func postorderTraversal(root *TreeNode) []int {
	var result []int
	stack := make([]*TreeNode, 0)
	cur := root
	stack = append(stack, cur)
	//原则:left先进栈，right后进栈
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		result = append(result, node.Val)
		stack = stack[:len(stack)-1] //出栈
		if node.Left != nil { //left先进栈
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

	}
	//反转result
	n := len(result)
	for i := 0; i < n/2; i++ {
		result[i], result[n-1-i] = result[n-1-i], result[i]
	}

	return result
}

/**
	  3
	 / \
	9  20
   / \ / \
  15 7 1
 */
func main() {
	// -----------二叉树-----------------
	node3 := &TreeNode{Val: 15}
	node4 := &TreeNode{Val: 7}
	node5 := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 9, Left: node3, Right: node4}
	node2 := &TreeNode{Val: 20, Left: node5}
	root := TreeNode{Val: 3, Left: node1, Right: node2}
	fmt.Println(postorderTraversal(&root))
}
