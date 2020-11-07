#### [59. 螺旋矩阵 II](https://leetcode-cn.com/problems/spiral-matrix-ii/)

给定一个正整数 *n*，生成一个包含 1 到 *n^2*所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

**示例:**

```
输入: 3
输出:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
```

---

关键代码，就是分步骤实现：

我们如果以坐标`（0,0）、（1,1）、（i,i）`为基准，然后分为四步：

- 第1步，行，左到右
- 第2步，列，上到下
- 第3步，行，右到左
- 第4步，列，下到上

注意，编程的时候，需要关注边界情况。

```go
for i := 0; i < base; i++ {
		end := n - i
		start := i
		//第1步，行，左到右
		for j := start; j < end; j++ {
			res[start][j] = num
			num++
		}
		//第2步，列，上到下
		for j := start + 1; j < end; j++ {
			res[j][end-1] = num
			num++
		}
		//第3步，行，右到左
		for j := end - 2; j >= start; j-- {
			res[end-1][j] = num
			num++
		}
		//第4步，列，下到上
		for j := end - 2; j > start; j-- {
			res[j][start] = num
			num++
		}
	}
```

