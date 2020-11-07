package main

import "fmt"

func setZeroes(matrix [][]int) {
	//需要原地改变
	n := len(matrix)
	if n == 0 {
		return
	}
	m := len(matrix[0])
	row := make(map[int]bool)
	column := make(map[int]bool)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == 0 {
				row[i] = true
				column[j] = true
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if row[i] || column[j] {
				matrix[i][j] = 0
			}
		}
	}
}

func main() {
	matrix := [][]int{
		{0, 0, 0, 5},
		{4, 3, 1, 4},
		{0, 1, 1, 4},
		{1, 2, 1, 3},
		{0, 0, 1, 1}}
	setZeroes(matrix)
	fmt.Println(matrix)
}
