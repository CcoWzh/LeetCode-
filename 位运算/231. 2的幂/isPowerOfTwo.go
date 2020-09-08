package main

import "fmt"

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}
	return n&(n-1) == 0
}

func main() {
	n := 16
	fmt.Println(isPowerOfTwo(n))
}
