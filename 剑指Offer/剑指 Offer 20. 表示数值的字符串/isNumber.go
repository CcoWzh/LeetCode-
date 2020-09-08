package main

import (
	"fmt"
	"strings"
)

/**
难点在于归纳各种正确的情况:
	‘.’出现正确情况：只出现一次，且在e的前面
	‘e’出现正确情况：只出现一次，且出现前有数字
	‘+’‘-’出现正确情况：只能在开头和e后一位
 */
func isNumber(s string) bool {
	//注意，需要预处理，去除s首位空格
	s = strings.TrimSpace(s)
	n := len(s)
	if n == 0 {
		return false
	}
	//3个标志位，判断对应的符号是否出现过
	pointFlag := 0
	eFlag := 0
	numFlag := 0
	for i := 0; i < n; i++ {
		cur := int(s[i] - '0')
		if cur >= 0 && cur <= 9 {
			numFlag = 1
		} else if s[i] == '.' && pointFlag == 0 && eFlag == 0 {
			//小数点只能出现一次，且只能出现在e前面
			pointFlag = 1
		} else if (s[i] == 'e' || s[i] == 'E') && eFlag == 0 && numFlag == 1 {
			//e前面必须得有数字，且只能出现一次
			eFlag = 1
			numFlag = 0
		} else if s[i] == '+' || s[i] == '-' {
			//正负号必须出现在开头或者e后面，且后面必须跟有数字，1e+
			if (numFlag == 0 && i == 0) || (eFlag == 1 && numFlag == 0 && i+1 != n) {
				continue
			} else {
				return false
			}
		} else {
			return false
		}
	}
	//返回的是numFlag！！！
	return numFlag == 1
}

func main() {
	s := "-1E-16"
	fmt.Println(isNumber(s))
}
