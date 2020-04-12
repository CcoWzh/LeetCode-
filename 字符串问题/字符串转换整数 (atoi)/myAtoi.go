package main

import "fmt"

//8. 字符串转换整数 (atoi)
func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	//去除字符串前的空格
	for i := 0; i < len(str); {
		if str[i]-'0' == 240 {
			i++
		} else {
			str = str[i:]
			break
		}
		if i == len(str)-1 {
			return 0
		}
	}
	var b int
	if str[0]-'0' == 251 || str[0]-'0' == 253 {
		b = int(str[0] - '0')-252
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
	if b == 253 {
		sum = 0 - sum
	}
	for i := 1; i < len(str); i++ {
		if int(str[i]-'0') > 9 {
			break
		}
		sum = sum * 10
		sum += int(str[i] - '0')
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
	str := "   -42"
	fmt.Println(myAtoi(str))
}
