#### [剑指 Offer 66. 构建乘积数组](https://leetcode-cn.com/problems/gou-jian-cheng-ji-shu-zu-lcof/)

给定一个数组 `A[0,1,…,n-1]`，请构建一个数组 `B[0,1,…,n-1]`，其中 `B` 中的元素 `B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]`。不能使用除法。

**示例:**

```
输入: [1,2,3,4,5]
输出: [120,60,40,30,24]
```

**提示：**

- 所有元素乘积之和不会溢出 32 位整数
- `a.length <= 100000`

---

暴力法：

```go
func constructArr(a []int) []int {
	n := len(a)
	res := make([]int, n)

	for i := 0; i < n; i++ {
		res[i] = 1
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			res[i] *= a[j]
		}
	}
	return res
}
```

遗憾的是，这个方法通过不了测试用例，会超出时间限制。

矩阵拆分：

```go
func constructArr(a []int) []int {
    if len(a) <= 1 {
        return a
    }
    left, right := make([]int, len(a)), make([]int, len(a))
    left[0] = 1
    right[len(a) - 1] = 1
    //左边到i的数的数组
    for i := 1; i < len(a); i++ {
        left[i] = left[i - 1] * a[i - 1]
    }
    //右边到i的数的数组
    for i := len(a) - 2; i >= 0; i-- {
        right[i] = right[i + 1] * a[i + 1]
    }
    //左右相乘，即为目标数组的值
    arr := make([]int, len(a))
    for i := 0; i< len(a); i++ {
        arr[i] = left[i] * right[i]
    }
    return arr
}
```

可将表格分为 **上三角** 和 **下三角** 两部分：

![Picture1.png](https://pic.leetcode-cn.com/6056c7a5009cb7a4674aab28505e598c502a7f7c60c45b9f19a8a64f31304745-Picture1.png)

##### 