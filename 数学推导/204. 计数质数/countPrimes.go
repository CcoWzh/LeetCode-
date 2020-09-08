package main

import "fmt"

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
		if mark[i] == 0 { //如果是素数，则他的倍数一定不是素数
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

func main() {
	n := 3
	fmt.Println(countPrimes(n))
}
