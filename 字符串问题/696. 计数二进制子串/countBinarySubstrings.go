package main

import "fmt"

func countBinarySubstrings(s string) int {
	n := len(s)
	container := make([]int, 0)
	result := 0
	for i := 0; i < n; {
		cur, j := s[i], 0
		for i < n && s[i] == cur {
			j++
			i++
		}
		container = append(container, j)
	}
	fmt.Println(container)

	for i := 0; i < len(container)-1; i++ {
		result += min(container[i], container[i+1])
	}

	return result
}

func min(a, b int) int {
	if a > b {
		a = b
	}
	return a
}

func main() {
	s := "1"
	fmt.Println(countBinarySubstrings(s))
}
