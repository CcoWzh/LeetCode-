package main

import (
	"fmt"
	"strconv"
)

//面试题 01.06.字符串压缩
func compressString(S string) string {
	if len(S) == 0 {
		return ""
	}
	result := string(S[0])
	left, right := 0, 0

	for right < len(S) {
		if S[right] != S[left] {
			result = result + strconv.Itoa(right-left) + string(S[right])
			left = right
		}
		right++
	}

	result = result + strconv.Itoa(right-left)
	if len(result) >= len(S) {
		return S
	}
	return result
}

func main() {
	S := "bb"
	fmt.Println(compressString(S))
}
