package main

import (
"fmt"
"strings"
)

/**
14. 最长公共前缀
 */
func longestCommonPrefix(strs []string) string {
	var s string
	if len(strs) == 0 {
		return ""
	}
	s = strs[0]
	for _, v := range strs {
		for {
			//返回s在v中首次出现的位置，没有的话返回-1
			if strings.Index(v, s) == 0 {
				break
			} else { //不断对s从后往前截取，直到符合条件
				s = s[:len(s)-1]
				if len(s) == 0 {
					return s
				}
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
	fmt.Println(strings.Index("flo", "fl"))
}
