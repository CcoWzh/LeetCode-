package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	str := []byte(s)
	for i, j := 0, n-1; i < j; {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}
	s = string(str)

	res := strings.Split(s, " ")
	for i, j := 0, len(res)-1; i < j; {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	s = ""
	for i := 0; i < len(res); i++ {
		if i == len(res)-1 {
			s += res[i]
			break
		}
		s += res[i] + " "
	}

	return s
}

func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))
}
