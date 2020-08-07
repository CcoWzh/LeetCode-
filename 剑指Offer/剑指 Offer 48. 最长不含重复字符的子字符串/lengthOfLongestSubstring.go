package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	//滑动窗口:[left,right]
	left, right := 0, 0
	//记录窗口里的值
	window := make(map[uint8]int)
	res := 0
	for right < len(s) {
		window[s[right]]++
		for window[s[right]] > 1 {
			window[s[left]]--
			left++
		}
		right++
		//因为我们要求的是最长子串，所以需要在每次移动 right 增大窗口时更新 res，
		//当要求的是最小值时，需要在移动 left 缩小窗口时更新
		res = max(res, right-left)
		//fmt.Println(s[left : right])
	}

	return res
}

func max(a, b int) int {
	if a < b {
		a = b
	}
	return a
}

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
