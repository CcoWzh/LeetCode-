package main

import (
	"fmt"
	"math"
)

func printNumbers(n int) []int {
	//由于n是正整数，所以不需要判断n小于0
	//1~(10^n-1)
	top := int(math.Pow(10, float64(n)))
	res := make([]int, top-1)

	for i := 1; i < top; i++ {
		res[i-1] = i
	}

	return res
}

func main() {
	n := 1
	fmt.Println(printNumbers(n))
}
