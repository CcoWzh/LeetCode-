#### [204. 计数质数](https://leetcode-cn.com/problems/count-primes/)

统计所有小于非负整数 *n* 的质数的数量。

**示例:**

```
输入: 10
输出: 4
解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
```

----

素数筛

```go
//小于非负整数 n 的质数
func countPrimes(n int) int {
	if n <= 1 {
		return 0
	}
	//素数筛
	mark := make([]int, n)
	mark[0], mark[1] = 1, 1
	for i := 1; i*i < n; i++ {
		//判断i是不是素数
		if isPrime(i) { //如果是素数，则他的倍数一定不是素数
			for j := 2 * i; j < n; j += i {
				mark[j] = 1
			}
		}
	}
	fmt.Println(mark)
	//统计所有是0的数，都是素数
	count := 0
	for i := 0; i < n; i++ {
		if mark[i] == 0 {
			count++
		}
	}

	return count
}

//判断是不是素数
func isPrime(n int) bool {
	if n == 2 {
		return true
	} else if n <= 1 {
		return false
	}
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
```

其实，都不需要判断一个数是不是素数：

```go
if mark[i] == 0 { //如果是素数，则他的倍数一定不是素数
			for j := 2 * i; j < n; j += i {
				mark[j] = 1
			}
		}
```

