package main

import "fmt"

//栈的表达方式，可以使用切片
func isValid(s string) bool {
	n := len(s)
	if n == 0 {
		return true
	}
	m := map[string]string{
		"}": "{",
		")": "(",
		"]": "["}
	stack := make([]string, 0)
	for i := 0; i < n; i++ {
		cur := string(s[i])
		if cur == "{" || cur == "[" || cur == "(" {
			stack = append(stack, cur)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != m[cur] {
				return false
			}
			//出栈
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	s := ""
	fmt.Println(isValid(s))
}
