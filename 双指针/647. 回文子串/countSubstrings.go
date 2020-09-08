package main

import "fmt"

func countSubstrings(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res += palindrome(s, i, i) +palindrome(s, i, i+1)
	}

	return res
}

//从中间扩展，判断回文
func palindrome(s string, left, right int) int {
	count := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
		count++
	}
	return count
}

func main() {
	s := "abc"
	fmt.Println(countSubstrings(s))
}
