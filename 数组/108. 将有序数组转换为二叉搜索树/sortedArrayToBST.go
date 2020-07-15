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
func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	root := &TreeNode{}
	mid := nums[n/2]
	root.Val = mid
	root.Left = help(nums[:n/2])
	root.Right = help(nums[n/2+1:])

	return root
}

func help(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	root := &TreeNode{}
	mid := nums[n/2]
	root.Val = mid
	root.Left = help(nums[:n/2])
	root.Right = help(nums[n/2+1:])
	return root
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	tree := sortedArrayToBST(nums)
	fmt.Println(inorderTraversal(tree))
}

//给定一个二叉树，返回它的中序 遍历。
//使用切片作为栈 stack = stack[:l-1]
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
		//fmt.Println(cur.Val)
		stack = stack[:len(stack)-1] // 出栈
		result = append(result, cur.Val)
		cur = cur.Right
	}
	return result
}

//简洁写法
//func sortedArrayToBST(nums []int) *TreeNode {
//	if len(nums) == 0 {
//		return nil
//	}
//
//	return &TreeNode{nums[len(nums)/2], sortedArrayToBST(nums[:len(nums)/2]), sortedArrayToBST(nums[len(nums)/2+1:])}
//}
