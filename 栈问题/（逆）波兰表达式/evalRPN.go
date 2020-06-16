package main

import (
	"fmt"
	"strconv"
)

/**
我们可以用一个栈S来实现计算，扫描从左往右进行，
如果扫描到操作数，则压进S，如果扫描到操作符，则从S弹出两个操作数

进行相应的操作，并将结果压进S(S的个数出2个进1个),当扫描结束后，S的栈顶就是表达式结果。
 */
func evalRPN(tokens []string) int {
	n := len(tokens)
	if n == 0 {
		return 0
	}

	stack := make([]int, 0)

	for i := 0; i < n; i++ {

		if tokens[i] != "+" && tokens[i] != "-" && tokens[i] != "*" && tokens[i] != "/" {
			ll, _ := strconv.Atoi(tokens[i])
			stack = append(stack, ll)
		} else {
			l := len(stack)
			//fmt.Println(stack)
			a := stack[l-1]
			b := stack[l-2]
			stack = stack[:l-2]
			if tokens[i] == "+" {
				stack = append(stack, a+b)
			} else if tokens[i] == "-" {
				stack = append(stack, b-a)
			} else if tokens[i] == "*" {
				stack = append(stack, b*a)
			} else if tokens[i] == "/" {
				stack = append(stack, b/a)
			}
		}
	}

	return stack[len(stack)-1]
}

func main() {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(tokens))
}
