package main

import (
	"fmt"
	"sort"
	"strings"
)

//820. 单词的压缩编码
func minimumLengthEncoding(words []string) int {
	var result int

	N := len(words)
	// 反转每个单词
	reversed_words := make([]string, len(words))
	for i := 0; i < N; i++ {
		word := words[i]
		rword := reverseString(word)
		reversed_words[i] = rword
	}
	// 字典序排序
	sort.Strings(reversed_words)
	fmt.Println(reversed_words)

	for i := 0; i < len(reversed_words); i++ {
		if i+1 < len(reversed_words) && strings.HasPrefix(reversed_words[i+1], reversed_words[i]) {
			continue
		} else {
			result += len(reversed_words[i]) + 1 // 单词加上一个 '#' 的长度
		}

	}

	return result
}

func reverseString(s string) string {
	runes := []rune(s)
	//runes := []byte(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}

func main() {
	words := []string{"time", "me", "bell"}
	fmt.Println(minimumLengthEncoding(words))
}
