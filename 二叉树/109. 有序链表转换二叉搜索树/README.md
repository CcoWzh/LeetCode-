#### [109. 有序链表转换二叉搜索树](https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/)

给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树*每个节点* 的左右两个子树的高度差的绝对值不超过 1。

**示例:**

```
给定的有序链表： [-10, -3, 0, 5, 9],

一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5
```

----

方法一：利用数组，转换成平衡二叉树

```go
func sortedListToBST(head *ListNode) *TreeNode {
	tree := make([]int, 0)
	for head != nil {
		tree = append(tree, head.Val)
		head = head.Next
	}
	//构建平衡树
	return buildTree(tree, 0, len(tree)-1)
}

func buildTree(tree []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := (start + end) / 2

	root := &TreeNode{Val: tree[mid]}
	root.Left = buildTree(tree, start, mid-1)
	root.Right = buildTree(tree, mid+1, end)

	return root
}
```

方法二：

更快就是直接获取头结点，先构建它呗，它是什么，它是 `BST` 最左子树的根节点啊。

什么遍历方式会优先遍历左子树，然后根节点，然后右子树？——中序遍历。

其实，`BST` 的中序遍历的节点次序，正好是该有序链表的节点次序。

因此，我们维护一个指针 h，让它去扫这个链表，它初始指向头节点，然后按中序遍历的次序，如下图，用 h 指向的节点值构建节点，构建完一个，h 指针就后移一位。

![image.png](https://pic.leetcode-cn.com/3016ccbd63d3b3d403583ee529e2d6f3e9a2b7d01ac4c45a9440d4b4685b9505-image.png)

即，求出中间节点，分成两部分，先根据左边部分递归构建左子树，这个递归就会首先创建上图中的 -10 节点，然后 h 指针后移，接着创建 -5 节点……直到构建完整个` BST`。

