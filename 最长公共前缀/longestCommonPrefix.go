package main

import (
	"fmt"
	"strings"
)

/**
最长公共前缀
 */
func longestCommonPrefix(strs []string) string {
	var s string
	if len(strs) == 0{
		return ""
	}
	s = strs[0]
	for _, v := range strs {
		for {
			if strings.Index(v, s) == 0 {
				break
			} else {
				s = s[:len(s)-1]
			}
		}
	}
	return s
}

func main() {
	strs1 := []string{"flower", "flow", "flight"}
	strs2 := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs1))
	fmt.Println(longestCommonPrefix(strs2))
	fmt.Println(strings.Index("flo", ""))
}
