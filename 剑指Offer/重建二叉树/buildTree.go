package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//var m map[int]int
//var po []int

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//func buildTree(preorder []int, inorder []int) *TreeNode {
//	po = preorder
//	m = make(map[int]int)
//	for i := 0; i < len(inorder); i++ {
//		m[inorder[i]] = i
//	}
//	return recur(0, 0, len(inorder)-1)
//}
//
//func recur(pre_root, in_left, in_right int) *TreeNode {
//	if in_left > in_right {
//		return nil
//	}
//	root := &TreeNode{Val: po[pre_root]}
//	i := m[po[pre_root]]
//	fmt.Println("left is : ", in_left)
//	fmt.Println("right is : ", in_right)
//	fmt.Println("left root is : ", pre_root+1)
//	fmt.Println("right root is : ", pre_root+i-in_left+1)
//	fmt.Println("==============================")
//	root.Left = recur(pre_root+1, in_left, i-1)
//	root.Right = recur(pre_root+i-in_left+1, i+1, in_right)
//	return root
//}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	for i, v := range inorder {
		if v == preorder[0] {
			root.Left = buildTree(preorder[1:i+1], inorder[:i])
			root.Right = buildTree(preorder[i+1:], inorder[i+1:])
		}
	}
	return root
}

func inorderTraversal(root *TreeNode) []int {
	var result []int
	stack := make([]*TreeNode, 0)
	cur := root
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur) //入栈
			cur = cur.Left
		}
		cur = stack[len(stack)-1:][0]
		stack = stack[:len(stack)-1] // 出栈
		result = append(result, cur.Val)
		cur = cur.Right
	}
	return result
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	root := buildTree(preorder, inorder)
	fmt.Println(inorderTraversal(root))
}
