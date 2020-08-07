package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	num1, num2 = reverseS(num1), reverseS(num2)
	n, m := len(num1), len(num2)
	i, j := 0, 0 //双指针
	sign := 0    //进位符
	res := ""
	for i < n || j < m { //两个数字位数不一定相同
		a, b := 0, 0
		//判断指针是否越界
		if i < n {
			a = int(num1[i] - '0')
		} else {
			a = 0
		}
		if j < m {
			b = int(num2[j] - '0')
		} else {
			b = 0
		}
		sum := a + b + sign
		//进位相加
		sign = sum / 10
		sum = sum - sign*10
		res += strconv.Itoa(sum)
		//指针相加
		i++
		j++
	}
	//如果最后还有进位的话，需要增加1
	if sign == 1 {
		res += "1"
	}

	return reverseS(res)
}

//反转字符串
func reverseS(s string) string {
	n, ss := len(s), []byte(s)
	if n == 0 {
		return ""
	}

	for i := 0; i < n/2; i++ {
		ss[i], ss[n-i-1] = ss[n-i-1], ss[i]
	}

	return string(ss)
}

func main() {
	num1 := "100000"
	num2 := "123456"
	fmt.Println(addStrings(num1, num2))
}
