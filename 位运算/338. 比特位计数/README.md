#### [338. 比特位计数](https://leetcode-cn.com/problems/counting-bits/)

给定一个非负整数 **num**。对于 **0 ≤ i ≤ num** 范围中的每个数字 **i** ，计算其二进制数中的 1 的数目并将它们作为数组返回。

**示例 1:**

```
输入: 2
输出: [0,1,1]
```

**示例 2:**

```
输入: 5
输出: [0,1,1,2,1,2]
```

**进阶:**

- 给出时间复杂度为**O(n\*sizeof(integer))**的解答非常容易。但你可以在线性时间**O(n)**内用一趟扫描做到吗？
- 要求算法的空间复杂度为**O(n)**。
- 你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的 **__builtin_popcount**）来执行此操作。

----

最简单的，就是暴力破解：

```go
func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 0; i <= num; i++ { //暴力遍历
		res[i] = bitNum(i)
	}
	return res
}
//计算数字n中有多少个1
func bitNum(n int) int {
	count := 0
	for n != 0 {
		n = n & (n - 1)
		count++
	}
	return count
}
```

利用奇偶性特点：

1. 如果 i 为偶数，那么`f(i) = f(i/2)`

   > 因为 i/2 本质上是i的二进制左移一位，低位补零，所以1的数量不变。

2. 如果 i 为奇数，那么`f(i) = f(i - 1) + 1`

   > 因为如果i为奇数，那么 i - 1必定为偶数，而偶数的二进制最低位一定是0，
   > 那么该偶数 +1 后最低位变为1且不会进位，所以奇数比它上一个偶数bit上多一个1，即 f(i) = f(i - 1) + 1。

代码：

```go
func countBits(num int) []int {
	res := make([]int, num+1)
	res[0] = 0
	for i := 0; i <= num; i++ {
		if i%2 == 0 { //i是偶数
			res[i] = res[i/2]
		} else { //i是奇数
			res[i] = res[i-1] + 1
		}
	}
	return res
}
```





​        
