package main

import (
	"fmt"
)

func multiply(num1 string, num2 string) string {
	n, m := len(num1), len(num2)
	if n == 0 || m == 0 {
		return ""
	}
	//考虑有0的情况
	if num1[0]-'0' == 0 || num2[0]-'0' == 0 {
		return "0"
	}
	product := make([]int, n+m)

	for i := n; i > 0; i-- {
		a := int(num1[i-1] - '0')
		for j := m; j > 0; j-- {
			b := int(num2[j-1] - '0')
			//最重要的就是这一步，检查了好久才发现了
			temp := a*b + product[i+j-1]
			product[i+j-1] = temp % 10
			product[i+j-2] += temp / 10
		}
	}
	s := ""
	for k, v := range product {
		if k == 0 && v == 0 {
			continue
		}
		s += string(v + '0') //这样编程更加简洁，不需要导入strconv.Itoa()函数
	}
	return s
}

func main() {
	num1 := "123"
	num2 := "456"
	fmt.Println(multiply(num1, num2))
}
