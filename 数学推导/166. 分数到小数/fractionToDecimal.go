package main

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	mark := 1
	if numerator*denominator < 0 {
		mark = 0
	}
	if numerator < 0 { //如果是负数的话，则先转换成正数
		numerator = -numerator
	}
	if denominator < 0 {
		denominator = -denominator
	} else if denominator == 0 {
		return ""
	}
	integer := strconv.Itoa(numerator / denominator) //整数部分
	remainder := numerator % denominator             //转换成真分子
	if mark == 0 { //注意分子和分母的符号
		integer = "-" + integer
	}

	//需要用一个哈希表记录余数出现在小数部分的位置，当你发现已经出现的余数，就可以将重复出现的小数部分用括号括起来。
	m := make(map[int]int)    //记录余数出现的位置
	decimal, cyclic := "", "" //decimal:小数部分,cyclic:循环的小数点
	count := 0
	for remainder != 0 {
		if index, ok := m[remainder]; ok { //如果有余数重复出现了，则代表有循环小数出现
			cyclic = decimal[index:]
			decimal = decimal[:index] + "(" + cyclic + ")"
			break
		}
		m[remainder] = count
		count++
		remainder *= 10
		decimal += strconv.Itoa(remainder / denominator)
		remainder = remainder % denominator
	}

	if len(decimal) == 0 {
		return integer
	}
	return integer + "." + decimal
}

func main() {
	numerator := 123   //分子
	denominator := 1 //分母
	fmt.Println(fractionToDecimal(numerator, denominator))
}
