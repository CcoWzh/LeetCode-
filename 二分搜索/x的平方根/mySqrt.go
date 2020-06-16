package main

import (
	"fmt"
	"math"
)

//牛顿法（有二分法）
//f(x)=x^2-a   ----> f(x)=f(x0)+(x-x0)f(xo)^-1
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	var cur float64
	cur = 1
	for {
		pre := cur
		cur = (cur + float64(x)/cur) / 2 //牛顿法
		if math.Abs(float64(cur-pre)) < 1e-6 {
			break
		}
		fmt.Println("cur is ", cur)
	}
	return int(cur)
}

func main() {
	fmt.Println(mySqrt(88))
}
