package main

import (
	"fmt"
	"strconv"
)

func findNthDigit(n int) int {
	//起始位置，位数，数位数量
	start, digit, count := 1, 1, 9

	for n > count {
		n = n - count
		digit++
		start = start * 10
		count = start * digit * 9
	}

	num := start + (n-1)/digit
	c := (n - 1) % digit
	return int(strconv.Itoa(num)[c] - '0')
}

func main() {
	n := 1000
	fmt.Println(findNthDigit(n))
}
