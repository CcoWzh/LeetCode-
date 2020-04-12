package main

import "fmt"

func sumNums(n int) int {
	if n <= 1 {
		return n
	}
	return n + sumNums(n-1)
}

func main() {
	fmt.Println(sumNums(4))
}
