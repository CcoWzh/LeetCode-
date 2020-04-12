#### [面试题26. 树的子结构](https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/)

输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)

B是A的子结构， 即 A中有出现和B相同的结构和节点值。

例如:
给定的树 A:

```go
     3
    / \
   4   5
  / \
 1   2
```


给定的树 B：

```go
   4 
  /
 1
```

返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

**示例 1：**

```go
输入：A = [1,2,3], B = [3,1]
输出：false
```

**示例 2：**

```go
输入：A = [3,4,5,1,2], B = [4,1]
输出：true
```

**限制：**

```go
0 <= 节点个数 <= 10000
```

----

注意：这里题目说了，`约定空树不是任意一个树的子结构`

比较优美的代码：


```go
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {
    if A==nil||B==nil{
        return false
    }
    if A.Val==B.Val{
        eq:=isEq(A,B)
        if eq{
            return eq
        }
    }
    return isSubStructure(A.Left,B)||isSubStructure(A.Right,B)
}
func isEq(a,b *TreeNode) bool{
    if b==nil{
        return true
    }
    if a==nil&&b!=nil{
        return false
    }
    if a.Val!=b.Val{
        return false
    }
    return isEq(a.Left,b.Left)&&isEq(a.Right,b.Right)
}
```

