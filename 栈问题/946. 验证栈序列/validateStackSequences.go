package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) != 0 && len(popped) != 0 && stack[len(stack)-1] == popped[0] {
			stack = stack[:len(stack)-1] //出栈
			popped = popped[1:]          //出队
			//fmt.Println(stack, popped)
		}
	}
	return len(stack) == 0 && len(popped) == 0
}

func main() {
	pushed := []int{1, 3, 5, 2, 7}
	popped := []int{1, 2, 7, 5, 3}
	fmt.Println(validateStackSequences(pushed, popped))
}

////比较优美的的编程方式
//func validateStackSequences(pushed []int, popped []int) bool {
//	stack := make([]int, 0)
//
//	i := 0
//	for _, push := range pushed {
//		stack = append(stack, push)
//		for len(stack) > 0 && i < len(popped) && stack[len(stack)-1] == popped[i] {
//			i++
//			stack = stack[:len(stack)-1]
//		}
//	}
//	return i == len(pushed)
//}
