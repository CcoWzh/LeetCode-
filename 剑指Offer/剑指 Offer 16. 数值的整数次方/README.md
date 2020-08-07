#### [剑指 Offer 16. 数值的整数次方](https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/)

实现函数double Power(double base, int exponent)，求base的exponent次方。不得使用库函数，同时不需要考虑大数问题。

**示例 1:**

```
输入: 2.00000, 10
输出: 1024.00000
```

**示例 2:**

```
输入: 2.10000, 3
输出: 9.26100
```

**示例 3:**

```
输入: 2.00000, -2
输出: 0.25000
解释: 2-2 = 1/22 = 1/4 = 0.25
```

**说明:**

- -100.0 < *x* < 100.0
- *n* 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。

注意：本题与主站 50 题相同：https://leetcode-cn.com/problems/powx-n/

----

计算相乘，这是普通版本的：

```go
//相乘,n>0
//这里可能会超时，可以使用快速幂
func compute(x float64, n int) float64 {
	q := float64(1)
	for i := 0; i < n; i++ {
		q *= x
	}

	return q
}
```

但是，这个很可能超时，所以，需要改进：

```go
//快速计算，n/2
func compute(x float64, n int) float64 {
	if n == 1 {
		return x
	} else if n == 0 {
		return 1
	}
	q := compute(x, n/2)
	q = q * q
	if n%2 == 1 {
		q *= x
	}

	return q
}
```

