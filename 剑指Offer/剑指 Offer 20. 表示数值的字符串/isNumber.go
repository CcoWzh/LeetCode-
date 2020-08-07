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
			//只出现一次(pointFlag == 0)，且在e的前面(eFlag == 0)
			pointFlag = 1
		} else if s[i] == 'e' && numFlag == 1 && eFlag == 0 {
			eFlag = 1
			//e后面一定要有数字(i+1 < n),12e
			numFlag = 0
		} else if s[i] == '+' || s[i] == '-' {
			//只能在开头和e后一位
			if (i != 0 && s[i-1] != 'e') || i+1 == n {
				//最后一位，不能是+-号，1e+
				return false
			}
		} else {
			return false
		}
	}
	//返回的是numFlag！！！，"."
	return numFlag == 1
}

func main() {
	s := " +1.2e-3 "
	fmt.Println(isNumber(s))
}
