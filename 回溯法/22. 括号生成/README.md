#### [22. 括号生成](https://leetcode-cn.com/problems/generate-parentheses/)

数字 *n* 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 **有效的** 括号组合。

**示例：**

```
输入：n = 3
输出：[
       "((()))",
       "(()())",
       "(())()",
       "()(())",
       "()()()"
     ]
```

----

回溯法

每次都只需要做两种选择：

```go
//l,r:当前左括号个数，当前右括号个数
func backTrace(n int, l, r int, path string, result *[]string) {
	if l+r == 2*n && l <= n && r <= n {
		*result = append(*result, path)
	}

	if l > r && l <= n {
		backTrace(n, l+1, r, path+"(", result)
		backTrace(n, l, r+1, path+")", result)
	} else if l <= r && l <= n {
		backTrace(n, l+1, r, path+"(", result)
	} else if l > n || r > n {
		return
	}
}
```

