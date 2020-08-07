package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	n := len(pushed)
	if n == 0 {
		return true
	}
	stack := make([]int, 0)
	//模拟入栈和出栈
	for i := 0; i < n; i++ {
		stack = append(stack, pushed[i]) //入栈
		for len(popped) != 0 && len(stack) != 0 && stack[len(stack)-1] == popped[0] {
			stack = stack[:len(stack)-1] //出栈
			popped = popped[1:]          //出栈，但是这更像是出队列
		}
	}
	//如果栈没有完全弹出，则说明不是弹出序列
	if len(popped) != 0 {
		return false
	}
	return true
}

func main() {
	pushed := []int{1, 2, 3, 4, 5}
	popped := []int{4, 5, 3, 1, 2}
	fmt.Println(validateStackSequences(pushed, popped))
}
