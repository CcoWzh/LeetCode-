#### [60. 第k个排列](https://leetcode-cn.com/problems/permutation-sequence/)

难度中等295收藏分享切换为英文关注反馈

给出集合 `[1,2,3,…,*n*]`，其所有元素共有 *n*! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 *n* = 3 时, 所有排列如下：

1. `"123"`
2. `"132"`
3. `"213"`
4. `"231"`
5. `"312"`
6. `"321"`

给定 *n* 和 *k*，返回第 *k* 个排列。

**说明：**

- 给定 *n* 的范围是 [1, 9]。
- 给定 *k* 的范围是[1,  *n*!]。

**示例 1:**

```
输入: n = 3, k = 3
输出: "213"
```

**示例 2:**

```
输入: n = 4, k = 9
输出: "2314"
```

---

通过找寻规律解题：

```go
var factorial = []int{0, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880, 3628800}

排列的第一位数字是： 
index := k / factorial[n-1]
result += nums[index]

之后，我们需要将数字 nums[index] 从当前的排列选择中删除
即，更新：
k = k % factorial[n-1]
n--

当发现k=0时，即以后的排列，都是按照 1,2,3,4,5...等正序排列的
```

常见的排列类问题有三类，即，相似题目：

- [46. 全排列](https://leetcode-cn.com/problems/permutations/)
- [31. 下一个排列](https://leetcode-cn.com/problems/next-permutation/)
- [60. 第k个排列](https://leetcode-cn.com/problems/permutation-sequence/)

