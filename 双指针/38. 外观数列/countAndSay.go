package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	str := countAndSay(n - 1)
	//str := "1211"
	res := ""
	l, r := 0, 0
	lenght := len(str)
	for r < lenght {
		curChar := str[l]
		for r < lenght && str[r] == curChar {
			r++
		}
		res += strconv.Itoa(r-l) + string(curChar)
		l = r
	}

	return res
}

func main() {
	n := 6
	fmt.Println(countAndSay(n))
}
