package main

import "fmt"

func constructArr(a []int) []int {
	n := len(a)
	if n == 0 {
		return nil
	}
	res := make([]int, n)
	//计算上半角和下半角的积
	res[0] = 1
	//i从正数第二个数字开始
	for i := 1; i < n; i++ {
		res[i] = a[i-1] * res[i-1]
	}
	temp := 1
	//i从倒数第二个数字开始
	for i := n - 2; i >= 0; i-- {
		temp *= a[i+1]
		res[i] *= temp
	}

	return res
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(constructArr(a))
}
