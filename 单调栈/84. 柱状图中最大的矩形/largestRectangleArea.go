package main

import "fmt"

func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}
	//单调栈，用于存储下标
	stack := make([]int, 0)
	stack = append(stack, -1) //设一个哨兵
	res := heights[0]

	for i := 0; i < n; i++ {
		cur := heights[i]
		//维护单调栈
		for len(stack) != 1 && heights[stack[len(stack)-1]] > cur {
			//在出栈的时候，维护最大值
			length := heights[stack[len(stack)-1]]
			//stack[len(stack)-2]是其左侧第一个比他小的元素（下标）
			//i 是其右侧第一个比他小的元素
			width := i - stack[len(stack)-2] - 1
			res = max(res, length*width)
			//出栈
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	//注意，上一步中单调栈可以还存在元素
	//需要对栈内元素计算，其宽度的计算是不一样的
	for len(stack) != 1 {
		length := heights[stack[len(stack)-1]]
		//对于栈顶元素，其右侧再没有比他跟小的元素了，所以是(n - 1)
		width := (n - 1) - stack[len(stack)-2]
		res = max(res, length*width)
		stack = stack[:len(stack)-1] //出栈
	}

	return res
}

func max(a, b int) int {
	if a < b {
		a = b
	}
	return a
}

func main() {
	heights := []int{6, 7, 5, 2, 4, 5, 9, 3}
	fmt.Println(largestRectangleArea(heights))
}
