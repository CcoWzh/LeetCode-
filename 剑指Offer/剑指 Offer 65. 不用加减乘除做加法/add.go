package main

import "fmt"

func add(a int, b int) int {
	n, c := a^b, (a&b)<<1

	for c != 0 {
		n = a ^ b
		c = (a & b) << 1
		a, b = n, c
	}

	return n
}

func main() {
	a, b := -3,9
	fmt.Println(add(a, b))
}
