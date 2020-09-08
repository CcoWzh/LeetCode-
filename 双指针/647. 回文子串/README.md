#### [647. 回文子串](https://leetcode-cn.com/problems/palindromic-substrings/)

给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。

具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。

**示例 1：**

```
输入："abc"
输出：3
解释：三个回文子串: "a", "b", "c"
```

**示例 2：**

```
输入："aaa"
输出：6
解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
```

**提示：**

- 输入的字符串长度不会超过 1000 。

---

双指针的使用

**寻找回文串的问题核心思想是：从中间开始向两边扩散来判断回文串**。

![img](https://mmbiz.qpic.cn/mmbiz_png/map09icNxZ4lLwdm05DtOeOPia4eSQF3HJ35jOicswr8BxewicbXvjKK3tpERQqORIqmJwddx7AXwxhjDm4QBicUoQw/640?wx_fmt=png&tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)

为什么要传入两个指针`l`和`r`呢？**因为这样实现可以同时处理回文串长度为奇数和偶数的情况**：

```go
for 0 <= i < len(s):
    # 找到以 s[i] 为中心的回文串
    palindrome(s, i, i)
    # 找到以 s[i] 和 s[i+1] 为中心的回文串
    palindrome(s, i, i + 1)
    更新答案
```

根据这个伪代码，直接暴力解决

参考：

- [经典面试题：最长回文子串](https://mp.weixin.qq.com/s/ux6VSWAPwgOS32xlScr2kQ)