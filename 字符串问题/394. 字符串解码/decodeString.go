package main

import (
	"fmt"
	"strconv"
	"strings"
)

func decodeString(s string) string {
	stack := make([]string, 0)

	index := 0
	for index < len(s) {
		cur := s[index]
		if cur >= '0' && cur <= '9' {
			digital := getDigital(s, &index)
			stack = append(stack, digital)
		} else if cur == '[' || (cur >= 'a' && cur <= 'z') || (cur >= 'A' && cur <= 'Z') {
			stack = append(stack, string(cur))
			index++
		} else if cur == ']' {
			//读取[]中的字符串
			temp := []string{}
			i := len(stack) - 1
			for stack[i][0] != '[' {
				temp = append(temp, stack[i])
				i--
			}
			//将[]中的字符串出栈
			stack = stack[:i]
			//反转字符串
			for i := 0; i < len(temp)/2; i++ {
				temp[i], temp[len(temp)-i-1] = temp[len(temp)-i-1], temp[i]
			}
			//将字符串数组转换成字符串
			ss := sliceToSring(temp)
			//将数字出栈
			count, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			//重复，恢复编码
			temp_s := strings.Repeat(ss,count)
			stack = append(stack, temp_s)
			//下标++
			index++
		}
	}

	return sliceToSring(stack)
}

//得到s，从per位置开始后的数字
func getDigital(s string, per *int) string {
	digital := ""

	for s[*per] >= '0' && s[*per] <= '9' && *per < len(s) {
		digital += string(s[*per])
		*per++
	}
	//fmt.Println(digital, *per)
	//x, _ := strconv.Atoi(digital)
	//fmt.Println(x)
	return digital
}

//将切片转换成string
func sliceToSring(s []string) string {
	result := ""
	for _,v := range s{
		result += v
	}
	return result
}

func main() {
	//s := "3[acd]2[bc]"
	s := "3[a2[c]]"
	fmt.Println(decodeString(s))
}
