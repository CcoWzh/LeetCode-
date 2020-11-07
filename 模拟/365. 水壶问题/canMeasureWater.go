package main

import "fmt"

//一开始理解错了题目的意思，以为可以不用限制中间过程中水放在哪儿。
//原来是有限制的，只能用这两个桶，不能借助其他桶
func canMeasureWater(x int, y int, z int) bool {
	if z == 0 {
		return true
	}
	if x+y < z {
		return false
	}

	if x < y {
		x, y = y, x
	}
	if y == 0 {
		return x == z
	}

	for x%y != 0 {
		x, y = y, x%y
	}
	//我们只需要找到 x, y的最大公约数并判断 z是否是它的倍数即可
	return z%y == 0
}

func main() {
	x, y, z := 3, 5, 4
	fmt.Println(canMeasureWater(x, y, z))
}
