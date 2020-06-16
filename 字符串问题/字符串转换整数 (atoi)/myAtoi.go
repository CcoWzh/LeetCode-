package main

import (
	"fmt"
	"strings"
)

//8. 字符串转换整数 (atoi)
func myAtoi(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0{
		return 0
	}

	var b int
	if str[0] == '+' || str[0] == '-' {
		b = int(str[0]-'0') - 252
		str = str[1:]
		if len(str) == 0 {
			return 0
		}
	}
	//第一个非空字符不是数字或正、负号。
	if str[0]-'0' > 9 {
		return 0
	}
	//计算字符串数字
	sum := int(str[0] - '0')
	//加符号
	if b == 1 {
		sum = 0 - sum
		b = -1
	} else {
		b = 1
	}
	for i := 1; i < len(str); i++ {
		if int(str[i]-'0') > 9 {
			break
		}
		sum = sum * 10
		sum = sum + b*int(str[i]-'0')
		//加限制条件
		if sum > 1<<31-1 {
			return 1<<31 - 1
		}
		if sum < -2147483648 {
			return -2147483648
		}
	}

	return sum
}

func main() {
	str := " "
	fmt.Println(myAtoi(str))
}
