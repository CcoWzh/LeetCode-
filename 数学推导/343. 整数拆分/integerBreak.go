package main

import "fmt"

//n=a*x ===> 当x相等时（几何算术不等式），最大 ===> y=(x^(1/x))^n
//y=e时，取最大值
func integerBreak(n int) int {
	//你可以假设 n 不小于 2 且不大于 58
	if n == 3 {
		return 2
	} else if n == 2 {
		return 1
	}

	q := n / 3
	re := n - 3*q
	//结尾是4的情况
	if re == 1 {
		q = q - 1
		re = 4
	}
	//求乘积
	p := 1
	for i := 0; i < q; i++ {
		p *= 3
	}
	if re != 0 {
		p *= re
	}

	return p
}

func main() {
	n := 9
	fmt.Println(integerBreak(n))
}
