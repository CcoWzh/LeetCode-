package main

import (
	"fmt"
)

//二分法
func mySqrt(x int) int {
	left, right := 0, x
	//由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。
	res := -1
	for left <= right {
		mid := left + (right-left)/2
		q := mid * mid
		if q <= x {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	//fmt.Println(left)
	return res
}

func main() {
	fmt.Println(mySqrt(4))
}
