package main

import "fmt"

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

func main() {
	num := 7
	fmt.Println(countBits(num))
}
