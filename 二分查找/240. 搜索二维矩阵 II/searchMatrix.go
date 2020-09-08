package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	row, column := 0, m-1

	for row < n && column >= 0 {
		cur := matrix[row][column]
		if cur == target {
			return true
		} else if cur > target { //说明，target不可能出现再数字cur所在的列
			column--
		} else {
			//说明，target在(row,column)的右边或者下边
			//但是，由于右边搜索过来，所以只可能出现在下边
			row++
		}
	}

	return false
}

//每行的元素从左到右升序排列。
//每列的元素从上到下升序排列。
func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30}}
	target := 9
	fmt.Println(searchMatrix(matrix, target))
}
