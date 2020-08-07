package main

import "fmt"

/**
实现 strStr() 函数。
给定一个 haystack 字符串和一个 needle 字符串，
在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。
 */
func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	//当 needle 是空字符串时，我们应当返回什么值呢
	if m == 0 {
		return 0
	}
	//似乎，这个也是可以使用指针的吧
	for i := 0; i < n; i++ {
		if haystack[i] == needle[0] && n-i >= m {
			//匹配
			temp := i
			j := 0
			for j < m {
				if haystack[temp] == needle[j] {
					temp++
					j++
				} else {
					break
				}
				fmt.Println(haystack[temp-1], needle[j-1])
			}

			if j == m {
				return i
			}

		}
	}

	return -1
}

func main() {
	haystack := ""
	needle := ""
	fmt.Println(strStr(haystack, needle))
}

//优雅的实现
//func strStr(haystack string, needle string) int {
//	if needle == "" {
//		return 0
//	}
//
//	l2 := len(needle)
//	for i := 0; i <= len(haystack)-l2; i++ {
//		if haystack[i:i+l2] == needle {
//			return i
//		}
//	}
//
//	return -1
//}
