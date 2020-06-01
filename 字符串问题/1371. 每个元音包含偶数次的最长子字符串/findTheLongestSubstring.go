package main

import "fmt"

func findTheLongestSubstring(s string) int {
	dp := make([]int, 32)
	for i := 0; i < 32; i++ {
		dp[i] = -1
	}

	pattern, result := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			pattern ^= 1 << 0
		} else if s[i] == 'e' {
			pattern ^= 1 << 1
		} else if s[i] == 'i' {
			pattern ^= 1 << 2
		} else if s[i] == 'o' {
			pattern ^= 1 << 3
		} else if s[i] == 'u' {
			pattern ^= 1 << 4
		}

		if dp[pattern] != -1 {
			//这一步，真是妙啊
			cur_len := i - dp[pattern]
			result = max(result, cur_len)
		} else {
			dp[pattern] = i
		}

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
	s := "eleetminicoworoep"
	fmt.Println(findTheLongestSubstring(s))
}
