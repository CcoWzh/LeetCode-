package main

import "fmt"

func isHappy(n int) bool {
	//如果一个数重复出现了，那就代表出现了循环嘛
	slow, fast := n, happy(n)

	for fast != 1 && slow != fast {
		fast = happy(happy(fast))
		slow = happy(slow)
	}

	return fast == 1
}

//返回平方和相加
func happy(n int) int {
	res := 0
	for n != 0 {
		res += (n % 10) * (n % 10)
		n = n / 10
	}
	return res
}

func main() {
	n := 116
	fmt.Println(isHappy(n))
}
