package main

import "fmt"

type TreeNode struct {
	Data   interface{}
	Lchild *TreeNode //左子树
	Rchild *TreeNode //右子树
}

//前序遍历：根 -----> 左子树 -----> 右子树
//递归实现
func preOrderTraverse(tree *TreeNode) {
	if tree == nil {
		return
	} else {
		fmt.Print(tree.Data, " ")
		preOrderTraverse(tree.Lchild)
		preOrderTraverse(tree.Rchild)
	}
}

//非递归前序遍历
func NpreOrderTraverse(tree *TreeNode) {

}

//中序遍历： 左 ---> 根 --->右
func inOrderTraverse(tree *TreeNode) {
	if tree == nil {
		return
	}
	inOrderTraverse(tree.Lchild)
	fmt.Print(tree.Data, " ")
	inOrderTraverse(tree.Rchild)
}

//后序遍历： 左 ----> 右 ---> 根
func postTraverse(tree *TreeNode) {
	if tree == nil {
		return
	} else {
		postTraverse(tree.Lchild)
		postTraverse(tree.Rchild)
		fmt.Print(tree.Data, " ")
	}
}

/**
	  3
	 / \
	9  20
   / \ /
  15 7 1
 */
func main() {
	// -----------二叉树-----------------
	node3 := &TreeNode{Data: 15}
	node4 := &TreeNode{Data: 7}
	node5 := &TreeNode{Data: 1}
	node1 := &TreeNode{Data: 9, Lchild: node3, Rchild: node4}
	node2 := &TreeNode{Data: 20, Lchild: node5}
	root := TreeNode{Data: 3, Lchild: node1, Rchild: node2}
	//root.Lchild = &TreeNode{Data: 9}
	//root.Rchild = &TreeNode{Data: 20}
	//root.Lchild.Lchild = &TreeNode{Data: 15}
	//root.Lchild.Rchild = &TreeNode{Data: 7}
	//root.Rchild.Lchild = &TreeNode{Data: 1}

	fmt.Print("前序遍历：")
	preOrderTraverse(&root)

	fmt.Print("\n中序遍历：")
	inOrderTraverse(&root)

	fmt.Print("\n后序遍历：")
	postTraverse(&root)
}
