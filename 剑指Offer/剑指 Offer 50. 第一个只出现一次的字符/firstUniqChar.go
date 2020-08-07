package main

import "fmt"

//cache :=[27]int{},用这个更佳
func firstUniqChar(s string) byte {
	n := len(s)
	if n == 0 {
		return byte(32)
	}
	m := make(map[uint8]int)
	for i := 0; i < n; i++ {
		m[s[i]]++
	}

	for i := 0; i < n; i++ {
		if m[s[i]] == 1 {
			return byte(s[i])
		}
	}

	return byte(32)
}

func main() {
	s := " "
	fmt.Println(firstUniqChar(s))
}
