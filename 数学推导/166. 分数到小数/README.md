#### [166. 分数到小数](https://leetcode-cn.com/problems/fraction-to-recurring-decimal/)

给定两个整数，分别表示分数的分子 numerator 和分母 denominator，以字符串形式返回小数。

如果小数部分为循环小数，则将循环的部分括在括号内。

**示例 1:**

```
输入: numerator = 1, denominator = 2
输出: "0.5"
```

**示例 2:**

```
输入: numerator = 2, denominator = 1
输出: "2"
```

**示例 3:**

```
输入: numerator = 2, denominator = 3
输出: "0.(6)"
```

---

思路：

[核心思想是当余数出现循环的时候，对应的商也会循环](https://leetcode-cn.com/problems/fraction-to-recurring-decimal/solution/fen-shu-dao-xiao-shu-by-leetcode/)。


算法：**需要用一个哈希表记录余数出现在小数部分的位置，当你发现已经出现的余数，就可以将重复出现的小数部分用括号括起来**。

这道题目本身并不难，但是需要把所以情况都考虑进去，就很困难了