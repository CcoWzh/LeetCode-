package main

import "fmt"

//70. 爬楼梯
func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	a, b := 1, 2
	result := 0
	for i := 3; i <= n; i++ {
		result = a + b
		a = b
		b = result
	}

	return result
}

func main() {
	n := 8
	fmt.Println(climbStairs(n))
}
