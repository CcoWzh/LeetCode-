#### [416. 分割等和子集](https://leetcode-cn.com/problems/partition-equal-subset-sum/)

给定一个**只包含正整数**的**非空**数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

**注意:**

1. 每个数组中的元素不会超过 100
2. 数组的大小不会超过 200

**示例 1:**

```
输入: [1, 5, 11, 5]

输出: true

解释: 数组可以分割成 [1, 5, 5] 和 [11].
```

 

**示例 2:**

```
输入: [1, 2, 3, 5]

输出: false

解释: 数组不能分割成两个元素和相等的子集.
```

---

0-1背包问题：

那么对于这个问题，我们可以先对集合求和，得出`sum`，把问题转化为背包问题：

**给一个可装载重量为`sum/2`的背包和`N`个物品，每个物品的重量为`nums[i]`。现在让你装物品，是否存在一种装法，能够恰好将背包装满**？



参考：

- [经典动态规划：0-1背包问题的变体](https://mp.weixin.qq.com/s/OzdkF30p5BHelCi6inAnNg)
- [小米笔试题（牛客网）](https://www.nowcoder.com/practice/7d78d8f671c2461aaeb304efb74b2310?tpId=125&&tqId=33757&rp=1&ru=/ta/exam-xiaomi&qru=/ta/exam-xiaomi/question-ranking)

