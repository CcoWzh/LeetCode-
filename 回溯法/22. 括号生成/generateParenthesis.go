package main

import "fmt"

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	backTrace(n, 0, 0, "", &result)
	return result
}

//l,r:当前左括号个数，当前右括号个数
func backTrace(n int, l, r int, path string, result *[]string) {
	if l+r == 2*n && l <= n && r <= n {
		*result = append(*result, path)
	}
	//每次都只需要做两种选择
	if l > r && l <= n {
		backTrace(n, l+1, r, path+"(", result)
		backTrace(n, l, r+1, path+")", result)
	} else if l <= r && l <= n {
		backTrace(n, l+1, r, path+"(", result)
	} else if l > n || r > n {
		return
	}
}

func main() {
	n := 0
	fmt.Println(generateParenthesis(n))
}
