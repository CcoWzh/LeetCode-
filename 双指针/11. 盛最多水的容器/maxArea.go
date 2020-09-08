package main

import "fmt"

func maxArea(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	left, right := 0, n-1
	res := 0
	for left < right {
		w := right - left                     //宽
		h := min(height[left], height[right]) //长
		s := w * h                            //面积
		if s > res { //计算最大值
			res = s
		}
		//移动指针
		if height[left] > height[right] { //如果右指针的值小，则移动右指针
			right--
		} else {
			left++
		}

	}
	return res
}

func min(a, b int) int {
	if a > b {
		a = b
	}
	return a
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
