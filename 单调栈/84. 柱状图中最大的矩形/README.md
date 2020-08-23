#### [84. 柱状图中最大的矩形](https://leetcode-cn.com/problems/largest-rectangle-in-histogram/)

给定 *n* 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。

 

![img](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/histogram.png)

以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 `[2,1,5,6,2,3]`。

 

![img](https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/12/histogram_area.png)

图中阴影部分为所能勾勒出的最大矩形面积，其面积为 `10` 个单位。

**示例:**

```
输入: [2,1,5,6,2,3]
输出: 10
```

---

单调递减栈

这里困难的，是需要确定下标，也就是宽度的选取：

```go
func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}
	//单调栈，用于存储下标
	stack := make([]int, 0)
	stack = append(stack, -1) //设一个哨兵
	res := heights[0]

	for i := 0; i < n; i++ {
		cur := heights[i]
		//维护单调栈
		for len(stack) != 1 && heights[stack[len(stack)-1]] > cur {
			//在出栈的时候，维护最大值
			length := heights[stack[len(stack)-1]]
			//stack[len(stack)-2]是其左侧第一个比他小的元素（下标）
			//i 是其右侧第一个比他小的元素
			width := i - stack[len(stack)-2] - 1
			res = max(res, length*width)
			//出栈
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	//注意，上一步中单调栈可以还存在元素
	//需要对栈内元素计算，其宽度的计算是不一样的
	for len(stack) != 1 {
		length := heights[stack[len(stack)-1]]
		//对于栈顶元素，其右侧再没有比他跟小的元素了，所以是(n - 1)
		width := (n - 1) - stack[len(stack)-2]
		res = max(res, length*width)
		stack = stack[:len(stack)-1] //出栈
	}

	return res
}

func max(a, b int) int {
	if a < b {
		a = b
	}
	return a
}
```

相关题目：

- [42. 接雨水（困难）](https://leetcode-cn.com/problems/trapping-rain-water/)

- [739. 每日温度（中等）](https://leetcode-cn.com/problems/daily-temperatures/)
- [496. 下一个更大元素 I（简单）](https://leetcode-cn.com/problems/next-greater-element-i/)
- [ 316. 去除重复字母（困难）](https://leetcode-cn.com/problems/remove-duplicate-letters/)
- [901. 股票价格跨度（中等）](https://leetcode-cn.com/problems/online-stock-span/)
- [ 402. 移掉K位数字](https://leetcode-cn.com/problems/remove-k-digits/)
- [581. 最短无序连续子数组](https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray/)