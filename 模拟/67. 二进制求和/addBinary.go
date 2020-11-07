package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	mark := 0
	res := ""
	index1, index2 := len(a)-1, len(b)-1

	for index1 >= 0 || index2 >= 0 {
		sum := 0
		x, y := 0, 0
		if index1 >= 0 {
			x = int(a[index1] - '0')
		}
		if index2 >= 0 {
			y = int(b[index2] - '0')
		}

		sum = x + y + mark
		mark = sum / 2
		sum = sum % 2
		res = strconv.Itoa(sum) + res
		index1--
		index2--
	}

	if mark > 0 {
		res = strconv.Itoa(mark) + res
	}
	//反转res
	return res
}

func main() {
	a := "10"
	b := "1011"
	fmt.Println(addBinary(a, b))
}
