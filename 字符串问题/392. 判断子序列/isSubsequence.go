package main

import "fmt"

func isSubsequence(s string, t string) bool {
	n, m := len(s), len(t)
	if n == 0 {
		return true
	} else if m < n {
		return false
	}
	//双指针
	i, j := 0, 0
	for i < m {
		if s[j] == t[i] {
			j++
			i++
		} else {
			i++
		}

		if j == n {
			break
		}
	}

	return j == n
}

func main() {
	s := ""
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}
