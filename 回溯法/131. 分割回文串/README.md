#### [131. 分割回文串](https://leetcode-cn.com/problems/palindrome-partitioning/)

给定一个字符串 *s*，将 *s* 分割成一些子串，使每个子串都是回文串。

返回 *s* 所有可能的分割方案。

**示例:**

```
输入: "aab"
输出:
[
  ["aa","b"],
  ["a","a","b"]
]
```

----

回溯法

需要注意的是：**需要复制path,要不然，他会改变，因为切片是浅拷贝**

```go
func backTrace(s string, path []string, res *[][]string) {
	if len(s) == 0 {
		tmp := make([]string, len(path))
		copy(tmp, path) //需要复制path,要不然，他会改变，因为切片是浅拷贝
		*res = append(*res, tmp)
	}
	temp := path
	for i := 1; i <= len(s); i++ {
		cur := s[:i]
		if isPalindrome(cur) {
			path = append(path, cur)
			backTrace(s[i:], path, res)
		} else {
			continue
		}
		path = temp
	}
}
```

