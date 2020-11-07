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
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > p.Val && root.Val > q.Val { //说明p,q都在左子树中
		return lowestCommonAncestor(root.Left, p, q)
	}
	if root.Val < p.Val && root.Val < q.Val { //说明p,q都在右子树中
		return lowestCommonAncestor(root.Right, p, q)
	}
	//以上都不是，说明p,q分布式root的两边
	return root
}

/**
	  6
	 / \
	2   8
   / \ / \
  0  4 7  9
 */
func main() {
	// -----------二叉树-----------------
	node3 := &TreeNode{Val: 0}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 7}
	node6 := &TreeNode{Val: 9}
	node1 := &TreeNode{Val: 2, Left: node3, Right: node4}
	node2 := &TreeNode{Val: 8, Left: node5, Right: node6}
	root := &TreeNode{Val: 6, Left: node1, Right: node2}
	fmt.Println(lowestCommonAncestor(root, node1, node4).Val)
}
