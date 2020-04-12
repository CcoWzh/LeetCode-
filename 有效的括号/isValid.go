package main

import "fmt"

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	m := make(map[string]interface{})
	m["}"] = "{"
	m[")"] = "("
	m["]"] = "["

	strck := createStack(len(s))

	for _, v := range s {
		fmt.Println(strck.top)

		if string(v) == "{" || string(v) == "(" || string(v) == "[" {
			strck.Push(string(v))
		} else {
			s, _ := strck.Pop()
			if m[string(v)] != s {
				return false
			}
		}
	}

	if strck.top == 0 {
		return true
	}
	return false
}

// 创建并初始化栈，返回strck
func createStack(size int) Stack {
	s := Stack{}
	s.size = size
	s.top = 0
	s.data = make([]interface{}, size)
	return s
}

// 栈结构体
type Stack struct {
	size int // 容量
	top  int // 栈顶
	// 用slice作容器，定义为interface{}
	data []interface{}
}

//入栈，栈顶升高
func (t *Stack) Push(element interface{}) bool {
	t.data[t.top] = element
	t.top++
	return true
}

//出栈，栈顶下降
func (t *Stack) Pop() (r interface{}, err error) {
	t.top--
	if t.top < 0 {
		return nil, err
	}

	return t.data[t.top], nil
}

func main() {
	fmt.Print(isValid("()]"))
}

//func isValid(s string) bool {
//	validMap := map[int32]int32{'}': '{', ']': '[', ')': '('} //1.括号key是反的一定要注意
//	stack := make([]int32, 0)
//	for _, v := range s { //if elif elif else
//		l := len(stack)
//		if v == '[' || v == '(' || v == '{' {
//			stack = append(stack, v)
//		} else if l == 0 {
//			return false
//		} else if vv, ok := validMap[v]; ok && vv == stack[l-1] {
//			stack = stack[:l-1]
//		} else {
//			return false
//		}
//	}
//	return len(stack) == 0
//}
