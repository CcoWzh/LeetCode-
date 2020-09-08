#### [221. 最大正方形](https://leetcode-cn.com/problems/maximal-square/)

在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。

**示例:**

```
输入: 

1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0

输出: 4
```

----

这个居然用动态规划？确实，如果使用深度或广度优先搜索的话，则代码量巨大，容易出错不说，关键是判断是否可以组成正方形

这题和[1277. 统计全为 1 的正方形子矩阵](https://leetcode-cn.com/problems/count-square-submatrices-with-all-ones/) ，非常相似

`dp[i][j]`表示以 `(i, j)` 为右下角的正方形的最大边长,则，状态转移方程也一样：

```
dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
```

