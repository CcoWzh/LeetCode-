package main

import "fmt"

func myPow(x float64, n int) float64 {
	//0的负数次方，是不合法的
	//当n<0时，只需要在最后计算倒数即可
	flag := 1
	if n < 0 {
		flag = -1
		n = -n
	}
	//n 取绝对值
	q := compute(x, n)

	if flag == -1 {
		q = float64(1) / q
	}

	return q
}

//相乘,n>0
//这里可能会超时，可以使用快速幂
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

func main() {
	x, n := 0.0, -3
	fmt.Println(myPow(x, n))
}
