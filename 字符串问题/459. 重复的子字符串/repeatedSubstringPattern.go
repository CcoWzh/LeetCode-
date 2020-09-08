package main

import (
	"fmt"
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	x := s
	s = s+s
	return strings.Contains(s[1:len(s)-1],x)
}

func main() {
	fmt.Println(repeatedSubstringPattern("aa"))
}
