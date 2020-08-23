package main

import "fmt"

func rangeBitwiseAnd(m int, n int) int {
	var count uint8
	//寻找二进制最长公共前缀
	for m < n {
		m = m >> 1
		n = n >> 1
		count++
	}
	return m << count
}

func main() {
	m, n := 5, 7
	fmt.Println(rangeBitwiseAnd(m, n))
}
