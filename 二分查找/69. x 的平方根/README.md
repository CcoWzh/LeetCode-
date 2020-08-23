#### [69. x 的平方根](https://leetcode-cn.com/problems/sqrtx/)

实现 int sqrt(int x) 函数

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

```tsx
示例 1:
输入: 4
输出: 2

示例 2:
输入: 8
输出: 2
说明: 8 的平方根是 2.82842..., 
由于返回类型是整数，小数部分将被舍去。
```

方法一：牛顿法

```go
//牛顿法（有二分法）
//f(x)=x^2-a   ----> f(x)=f(x0)+(x-x0)f(xo)^-1
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	var cur float64
	cur = 1
	for {
		pre := cur
		cur = (cur + float64(x)/cur) / 2 //牛顿法
		if math.Abs(float64(cur-pre)) < 1e-6 {
			break
		}
		fmt.Println("cur is ", cur)
	}
	return int(cur)
}
```

方法二：二分法

```go
func mySqrt(x int) int {
	left, right := 0, x
	//由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
	res := -1
	for left <= right {
		mid := left + (right-left)/2
		q := mid * mid
		if q <= x {
			res = mid //更新
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	//fmt.Println(left)
	return res
}
```

一道简单的题，我在实际编程中，居然二分法重重出错:

首先，是`left = mid + 1`写成了`left++`,导致时间超时

再次是`res = mid`写成了`res = left`这样的低级错误

---

进阶：要求开平方的数，精确到小数点后4位

则，我们先不使用二分法，假设d的值不大于`1000000`,则，我们一个一个的遍历过去

```go
//开平方，精确到小数点
func getSqrt(d float64) float64 {
	n := 1.0
	for i := 0.0; i < 1000; {
		s := i * i
		if s == d || abs(s-d) < 0.001 {
			return i
		}
		if i*i < d { //当i的平方大于目标值时，加1，或者加0.1，0.01，0.001等等
			i += n
		}
		if i*i > d { //当i的平方小于目标值时,将跨度缩小10倍
			i = i - n
			n = n / 10
			i = i + n
		}
	}
	return 0
}

func abs(a float64) float64 {
	if a < 0 {
		a = -a
	}
	return a
}
```

