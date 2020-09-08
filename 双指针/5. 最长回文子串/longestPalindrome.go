package main

import "fmt"

/**
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
动态规划：最长公共子串+找下标筛选
// */
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	} //极限条件
	str := reverseString(s)
	var max, index int
	arr := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		arr[i] = make([]int, len(s))
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(str); j++ {

			if s[i] == str[j] {
				if i == 0 || j == 0 {
					arr[i][j] = 1
				} else {
					arr[i][j] = arr[i-1][j-1] + 1
				}
			}

			if arr[i][j] > max {
				if len(s)-1+arr[i][j]-1 == j+i {
					max = arr[i][j]
					index = i
				}
			}
		}
	}

	return s[index-max+1 : index+1]
}

/**
反转字符串
 */
func reverseString(s string) string {
	str := []byte(s)
	for i, j := 0, len(str)-1; i < len(str)/2; i++ {
		str[i], str[j] = str[j], str[i]
		j--
	}
	return string(str)
}

func main() {
	s := ""
	fmt.Println(longestPalindrome(s))
}
