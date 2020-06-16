package main

import "fmt"

var result []int

func spiralOrder(matrix [][]int) []int {
	rows  := len(matrix)
	result = []int{}

	if rows == 0{
		return result
	}

	columns := len(matrix[0])


	start := 0
	for columns > start*2 && rows > start*2 {
		printNums(matrix, start, rows, columns)
		start++
	}

	return result
}

func printNums(matrix [][]int, start, rows, columns int) {
	end_column := columns - 1 - start
	end_row := rows - 1 - start

	//从左到右打印
	for i := start; i <= end_column; i++ {
		result = append(result, matrix[start][i])
	}

	//从上到下打印
	if end_row > start {
		for i := start + 1; i <= end_row; i++ {
			result = append(result, matrix[i][end_column])
		}
	}

	//从右到左打印
	if end_row > start && end_column > start {
		for i := end_column - 1; i >= start; i-- {
			result = append(result, matrix[end_row][i])
		}
	}

	//从下到上打印
	if end_row-1 > start && end_column > start {
		for i := end_row - 1; i > start; i-- {
			result = append(result, matrix[i][start])
		}
	}

}

func main() {
	matrix := [][]int{}
	fmt.Println(spiralOrder(matrix))
}
