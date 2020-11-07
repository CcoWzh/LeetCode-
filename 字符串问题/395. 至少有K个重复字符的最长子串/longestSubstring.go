package main

import (
	"fmt"
)

func longestSubstring(s string, k int) int {
	charNum := make(map[rune]int) // 字符计数
	for _, v := range s {
		charNum[v]++
	}
	var split []int
	// 找到所有不满足条件的字符作为分割符，进行分治
	for index, v := range s {
		if charNum[v] < k {
			split = append(split, index)
		}
	}
	//如果整个字符串都符合要求，则直接返回这个字符串的长度
	if len(split) == 0 {
		return len(s)
	}
	split = append(split, len(s)) // 把最后的右边加进去
	result, left := 0, 0          // 结果，左起索引
	for _, v := range split {
		// 总长度不超过已知最长的就没必要比较了
		if v-left+1 <= result {
			continue
		}
		result = max(result, longestSubstring(s[left:v], k))
		left = v + 1
	}
	return result
}

func max(a, b int) int {
	if a < b {
		a = b
	}
	return a
}

func main() {
	s := "abcabb"
	k := 2
	fmt.Println(longestSubstring(s, k))
}
