package main

import "fmt"

//1071. 字符串的最大公因子
func gcdOfStrings(str1 string, str2 string) string {
	if str2+str1 != str1+str2 {
		return ""
	}
	a, b := len(str1), len(str2)
	for {
		c := a % b
		a = b
		b = c
		if c == 0 {
			break
		}
	}
	return str1[:a]
}

func main() {
	str2 := "ABABAB"
	str1 := "ABAB"
	fmt.Println(gcdOfStrings(str1, str2))

}
