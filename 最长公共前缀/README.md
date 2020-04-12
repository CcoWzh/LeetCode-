## 最长公共前缀

> 编写一个函数来查找字符串数组中的最长公共前缀。如果不存在公共前缀，则返回""

```tsx
示例 1:
输入: ["flower","flow","flight"]
输出: "fl"

示例 2:
输入: ["dog","racecar","car"]
输出: ""

解释: 输入不存在公共前缀。
说明:
所有输入只包含小写字母 a-z 。
```

#### 好的编程方案

```tsx
 1 func longestCommonPrefix(strs []string) string {
 2    if len(strs) < 1 {
 3        return ""
 4    }
 5    prefix := strs[0]
 6    for _,k := range strs {
 7        for strings.Index(k,prefix) != 0 {
 8            if len(prefix) == 0 {
 9                return ""
10            }
11            prefix = prefix[:len(prefix) - 1]
12        }
13    }
14    return prefix
15 }   
```

