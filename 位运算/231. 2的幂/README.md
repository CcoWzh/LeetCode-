#### [231. 2的幂](https://leetcode-cn.com/problems/power-of-two/)

给定一个整数，编写一个函数来判断它是否是 2 的幂次方。

**示例 1:**

```
输入: 1
输出: true
解释: 20 = 1
```

**示例 2:**

```
输入: 16
输出: true
解释: 24 = 16
```

**示例 3:**

```
输入: 218
输出: false
```

----

位运算：

```go
func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}
	//n&(n-1):消除数字 n 的二进制表示中的最后一个 1
	return n&(n-1) == 0
}
```

