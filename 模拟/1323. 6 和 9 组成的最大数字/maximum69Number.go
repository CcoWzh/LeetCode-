package main

import "fmt"

/**
1 <= num <= 10^4
num 每一位上的数字都是 6 或者 9 。
 */
func maximum69Number(num int) int {
	//num 最大是9999
	if num/1000 == 6 { //如果第一个数字是6，则改成9，返回
		num += 3000
	} else if (num%1000)/100 == 6 {
		num += 300
	} else if (num%100)/10 == 6 {
		num += 30
	} else if num%10 == 6 {
		num += 3
	}
	return num
}

func main() {
	num := 6696
	fmt.Println(maximum69Number(num))
}
