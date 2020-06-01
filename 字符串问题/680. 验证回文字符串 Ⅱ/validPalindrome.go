/**
给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
 */
package main

import "fmt"

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left <= right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			s = s[left : right+1]
			if palindrome(s[1:]) || palindrome(s[:len(s)-1]) {
				return true
			}
			break
		}
	}

	if left > right {
		return true
	} else {
		return false
	}
}

//双指针法
//判断是否是回文字符串
func palindrome(s string) bool {
	left, right := 0, len(s)-1

	for left <= right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			break
		}
	}

	if left > right {
		return true
	} else {
		return false
	}
}

func main() {
	s := ""
	fmt.Println(validPalindrome(s))
}
