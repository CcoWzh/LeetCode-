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
func flatten(root *TreeNode) {
	list := make([]*TreeNode, 0)
	//前序遍历，得到前序遍历的节点列表
	preorder(root, &list)
	//开始转换
	for i := 0; i < len(list)-1; i++ {
		cur, next := list[i], list[i+1]
		cur.Left, cur.Right = nil, next
		//fmt.Println(list[i].Val)
	}
}

func preorder(root *TreeNode, list *[]*TreeNode) {
	if root == nil {
		return
	} else {
		*list = append(*list, root)
		preorder(root.Left, list)
		preorder(root.Right, list)
	}
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
	flatten(&root)
	preOrderTraverse(&root)
}

//前序遍历：根 -----> 左子树 -----> 右子树
//递归实现
func preOrderTraverse(tree *TreeNode) {
	if tree == nil {
		return
	} else {
		fmt.Print(tree.Val, " ")
		preOrderTraverse(tree.Left)
		preOrderTraverse(tree.Right)
	}
}
