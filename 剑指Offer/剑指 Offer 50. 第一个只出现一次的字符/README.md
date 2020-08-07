#### [剑指 Offer 50. 第一个只出现一次的字符](https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/)

在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

**示例:**

```
s = "abaccdeff"
返回 "b"

s = "" 
返回 " "
```

**限制：**

```
0 <= s 的长度 <= 50000
```

---

更好的代码：

```go
func firstUniqChar(s string) byte {
    cache :=[27]int{} //不需要使用map
    for i:=0;i<len(s);i++{
        cache[s[i]-'a']++
    }
    for i:=0;i<len(s);i++{
        if cache[s[i]-'a'] == 1 {
            return s[i]
        }
    }
    return ' '
}
```

