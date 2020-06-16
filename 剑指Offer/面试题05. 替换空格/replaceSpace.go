package main

import "fmt"

func replaceSpace(s string) string {

	result := make([]string, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			result = append(result, "%20")
		} else {
			result = append(result, string(s[i]))
		}
	}

	ss := ""
	for i:=0;i< len(result);i++{
		ss += result[i]
	}

	return ss
}

func main() {
	s := "We are happy."
	fmt.Println(replaceSpace(s))
}
