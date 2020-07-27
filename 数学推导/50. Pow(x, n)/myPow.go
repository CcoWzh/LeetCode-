package main

import "fmt"

func myPow(x float64, n int) float64 {
	result := compute(x, abs(n))
	if n < 0 {
		result = 1 / result
	}
	return result
}

//求绝对值
func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

//计算
/**
if n%2 == 0 {
		result = compute(x, n/2) * compute(x, n/2)
	} else {
		result = compute(x, (n-1)/2) * compute(x, (n-1)/2) * x
	}
 */
func compute(x float64, n int) float64 {
	var result float64
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	//除以2
	result = compute(x, n>>1)
	result *= result
	//判断n是否为奇数
	if n&0x1 == 1 {
		result *= x
	}
	return result
}

func main() {
	fmt.Println(myPow(0.00001, 2147483647))
}
