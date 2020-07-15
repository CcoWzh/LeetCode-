package main

import "fmt"

func numTrees(n int) int {
	c := 1
	for i := 0; i < n; i++ {
		c = c * 2 * (2*i + 1) / (i + 2)
	}
	return c
}

func main() {
	n := 3
	fmt.Println(numTrees(n))
}
