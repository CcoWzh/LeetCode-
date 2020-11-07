#### [216. 组合总和 III](https://leetcode-cn.com/problems/combination-sum-iii/)

找出所有相加之和为 ***n*** 的 ***k\*** 个数的组合***。\***组合中只允许含有 1 - 9 的正整数，并且每种组合中不存在重复的数字。

**说明：**

- 所有数字都是正整数。
- 解集不能包含重复的组合。 

**示例 1:**

```
输入: k = 3, n = 7
输出: [[1,2,4]]
```

**示例 2:**

```
输入: k = 3, n = 9
输出: [[1,2,6], [1,3,5], [2,3,4]]
```

----

这里需要注意的是：

1. 在go语言中，需要使用`copy`来保证回溯法的撤销选择
2. 利用下标`start`，做到无重复

```go
func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	backTrace(0, n, k, 1, []int{}, &res)
	return res
}

func backTrace(cur, n, k int, start int, path []int, res *[][]int) {
	if cur == n && len(path) == k {
		buf := make([]int, len(path))
		copy(buf, path)
		*res = append(*res, path)
		return
	}
	if len(path) > k || cur > n {
		return
	}

	base := make([]int, len(path))
	copy(base, path)
	for i := start; i <= 9; i++ {
		if cur+i > n {
			break
		}
		path = append(path, i)
		backTrace(cur+i, n, k, i+1, path, res)
		path = base
	}
}
```

