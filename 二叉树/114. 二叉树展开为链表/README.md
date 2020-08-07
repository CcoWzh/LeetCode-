#### [114. 二叉树展开为链表](https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/)

给定一个二叉树，[原地](https://baike.baidu.com/item/原地算法/8010757)将它展开为一个单链表。

例如，给定二叉树

```
    1
   / \
  2   5
 / \   \
3   4   6
```

将其展开为：

```
1
 \
  2
   \
    3
     \
      4
       \
        5
         \
          6
```

---

前序遍历

第一反应，是看不懂什么是原地展开？原来就是：

```go
cur, next := list[i], list[i+1]
cur.Left, cur.Right = nil, next
```

也就是，在原来的节点上，改变原来的结构。树===>链表，就是其中一个子节点指向下一个节点，其余的子节点为空。