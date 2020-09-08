package main

import "fmt"

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
