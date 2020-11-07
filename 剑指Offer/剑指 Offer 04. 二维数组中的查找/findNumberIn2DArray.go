package main

/**
时间执行过长，没有明白题目中的提示（升序）
 */
func findNumberIn2DArray(matrix [][]int, target int) bool {
	row := len(matrix)
	if row == 0 {
		return false
	}
	column := len(matrix[0])

	i, j := 0, column-1
	for i < row && j >= 0 {
		cur := matrix[i][j]
		if cur == target {
			return true
		} else if cur < target {
			i++
		} else {
			j--
		}
	}

	return false
}

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	println(findNumberIn2DArray(matrix, 100))
}
