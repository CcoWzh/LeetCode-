package main

import (
	"fmt"
	"strconv"
)

func translateNum(num int) int {
	l := strconv.Itoa(num)
	//f(i) = f(i+1)+g*f(i+2)  **g=0||1
	return find(0, l)
}

func find(i int, num string) int {
	x, _ := strconv.Atoi(num[i:])
	if x <= 25 && x >= 10 {
		return 2
	} else if 0 <= x && x < 10 {
		return 1
	}

	cur := num[i] - '0'
	next := uint8(0)
	if i+1 >= len(num) {
		next = 0
	} else {
		next = num[i+1] - '0'
	}

	g := 0
	if cur*10+next <= 25 && cur*10+next >= 10 {
		g = 1
	}

	return find(i+1, num) + g*find(i+2, num)
}

func main() {
	num := 1225
	fmt.Println(translateNum(num))
	//fmt.Println(1 << 31)
}
