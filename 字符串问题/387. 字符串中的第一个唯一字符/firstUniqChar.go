package main

import "fmt"

func firstUniqChar(s string) int {
	word := [26]int{}
	n := len(s)
	for i := 0; i < n; i++ {
		word[s[i]-'a']++
	}

	for i := 0; i < n; i++ {
		if word[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	s := "q"
	fmt.Println(firstUniqChar(s))
}
