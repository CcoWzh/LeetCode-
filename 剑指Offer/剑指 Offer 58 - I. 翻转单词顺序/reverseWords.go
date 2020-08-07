package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	//s = strings.TrimSpace(s)  //这个其实可以不用
	str := strings.Fields(s)
	//反转字符串
	n := len(str)
	for i := 0; i < n/2; i++ {
		str[i], str[n-i-1] = str[n-i-1], str[i]
	}
	//拼接字符串
	//strings.Join(str, " ")
	res := ""
	for i := 0; i < n; i++ {
		if i == n-1 {
			res += str[i]
			break
		}
		res = res + str[i] + " "
	}
	return res
}

func main() {
	s := "  the good   boy    ha ha ha..  "
	fmt.Println(reverseWords(s))
}
