#### [剑指 Offer 32 - III. 从上到下打印二叉树 III](https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/)

请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。

 

例如:
给定二叉树: `[3,9,20,null,null,15,7]`,

```
    3
   / \
  9  20
    /  \
   15   7
```

返回其层次遍历结果：

```
[
  [3],
  [20,9],
  [15,7]
]
```

**提示：**

1. `节点总数 <= 1000`

----

```go
for i := 0; i < n; i++ {
			r := queue[i]
			//其实，广度遍历的顺序不需要改变，只需要改变承接的值的顺序
			if flag%2 == 0 {
				s := queue[n-i-1]
				tempRes = append(tempRes, s.Val)
			} else {
				tempRes = append(tempRes, r.Val)
			}
			if r.Left != nil {
				temp = append(temp, r.Left)
			}
			if r.Right != nil {
				temp = append(temp, r.Right)
			}
		}
```

