package main

import "fmt"

/**
限制：
1 <= n < s.length <= 10000
 */
func reverseLeftWords(s string, n int) string {
	for i := 0; i < n; i++ {
		s += string(s[i])
	}
	return s[n:]
}

func main() {
	s := "abcdefg"
	n := 6
	fmt.Println(reverseLeftWords(s, n))
}

/** 这么骚气的代码
func reverseLeftWords(s string, n int) string {
    res := s[n:] + s[:n]
	return res
}
 */
