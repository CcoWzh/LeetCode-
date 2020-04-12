package main

import "fmt"

//答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
//func fib(n int) int {
//	if n == 1 || n == 0 {
//		return n
//	}
//	a, b := 0, 1
//	for i := 0; i < n; i++ {
//		a, b = b, a+b
//		if a < 1e9+7 {
//			a = a % (1e9 + 7)
//			b = b % (1e9 + 7)
//		}
//	}
//	fmt.Println(a, b)
//	//45 溢出a
//	return a
//}

//青蛙跳台阶问题
func numWays(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	a, b := 1, 2
	for i := 1; i < n; i++ {
		a, b = b, a+b
		if a < 1e9+7 {
			a = a % (1e9 + 7)
			b = b % (1e9 + 7)
		}
	}
	fmt.Println(a, b)
	//45 溢出a
	return a
}

func main() {
	fmt.Println(numWays(2))
}
