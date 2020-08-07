package main

import "fmt"

//面试题62. 剑指 Offer 62. 圆圈中最后剩下的数字
//约瑟夫环问题
//f(n, m)表示n个数，m数的话，剩余的数字编号
func lastRemaining(n int, m int) int {
	if n == 0 {
		return 0
	}
	return (lastRemaining(n-1, m)+m)%n
}

func main() {
	n, m := 1,3
	fmt.Println(lastRemaining(n, m))
}
