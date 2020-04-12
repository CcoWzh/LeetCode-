package main

import "fmt"

//836. 矩形重叠
//将矩形重叠问题转化为区间重叠问题
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	x := !(rec1[2] <= rec2[0] || rec1[0] >= rec2[2])
	y := !(rec1[1] >= rec2[3] || rec1[3] <= rec2[1])
	return x && y
}

func main() {
	rec1 := []int{0, 0, 2, 2}
	rec2 := []int{1, 1, 3, 3}
	fmt.Println(isRectangleOverlap(rec1, rec2))
}
