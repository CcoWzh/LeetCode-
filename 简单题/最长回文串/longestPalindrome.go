package main

import "fmt"

//409. 最长回文串
func longestPalindrome(s string) int {
	var aCount [52]int //数组统计字符串个数
	for _, char := range s {
		if char-'a' >= 0 {
			aCount[char-'a'] ++
		} else {
			aCount[char-'A'+26] ++
		}
	}
	result, odd := 0, 0
	for _, v := range aCount {
		result += v
		if v%2 == 1 {
			odd++
		}
	}

	if odd != 0 {
		result = result - odd + 1 //减去奇数的个数，+1
	}
	return result
}

func main() {
	s := "abccccdd"
	fmt.Println(longestPalindrome(s))
}
