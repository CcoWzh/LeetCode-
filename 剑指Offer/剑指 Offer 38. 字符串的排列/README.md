#### [剑指 Offer 38. 字符串的排列](https://leetcode-cn.com/problems/zi-fu-chuan-de-pai-lie-lcof/)

输入一个字符串，打印出该字符串中字符的所有排列。

你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。

**示例:**

```
输入：s = "abc"
输出：["abc","acb","bac","bca","cab","cba"]
```

**限制：**

```
1 <= s 的长度 <= 8
```

---

这里需要注意的是，如果有重复的元素，怎么办？如，`aacc`

所以，在回溯的时候，需要先排序，再跳过：

```go
for i := 0; i < len(s); i++ {
		temp := s[i]
		fmt.Println(string(temp))
		//需要跳过重复的字符
		j := i
		for j+1 < len(s) && s[j+1] == temp {
			i++
			j++
		}
	}
```

然后，就是回溯法的一般套路了。