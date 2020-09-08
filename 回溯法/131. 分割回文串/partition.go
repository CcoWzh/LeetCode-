package main

import "fmt"

func partition(s string) [][]string {
	res := make([][]string, 0)
	backTrace(s, []string{}, &res)
	return res
}

func backTrace(s string, path []string, res *[][]string) {
	if len(s) == 0 {
		tmp := make([]string, len(path))
		copy(tmp, path) //需要复制path,要不然，他会改变，因为切片是浅拷贝
		*res = append(*res, tmp)
	}
	temp := path
	for i := 1; i <= len(s); i++ {
		cur := s[:i]
		if isPalindrome(cur) {
			path = append(path, cur)
			backTrace(s[i:], path, res)
		} else {
			continue
		}
		path = temp
	}
}

func isPalindrome(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	s := "aba"
	fmt.Println(partition(s))
	//fmt.Println(isPalindrome("aabca"))
}
