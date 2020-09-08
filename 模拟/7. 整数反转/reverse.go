package main

import "fmt"

func reverse(x int) int {
	res := 0
	for x != 0 {
		//这一步很重要，完美的避开了末尾是0的情况
		res = res*10 + x%10
		x = x / 10
	}
	//其数值范围为 [−2^31,  2^31 − 1]
	if res > 1<<31-1 || res < -(1 << 31) {
		return 0
	}
	return res
}

func main() {
	x := 120
	fmt.Println(reverse(x))
}
