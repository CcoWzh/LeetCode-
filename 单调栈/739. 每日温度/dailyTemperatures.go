package main

import "fmt"

func dailyTemperatures(T []int) []int {
	n := len(T)
	if n == 0 {
		return nil
	}
	stack := make([]int, 0)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		cur := T[i]
		for len(stack) != 0 && T[stack[len(stack)-1]] < cur { //单调栈
			topIndex := stack[len(stack)-1] //当前的数字大于栈顶元素时，说明就是第一个大于栈顶的元素
			stack = stack[:len(stack)-1]    //出栈
			res[topIndex] = i - topIndex    //当前index-栈顶的index，即是
		}
		stack = append(stack, i)
	}
	return res
}

func main() {
	T := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(T))
}
