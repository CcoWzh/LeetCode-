package main

import "fmt"

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	left, right := matrix[0][0], matrix[n-1][n-1]

	for left < right {
		mid := left + (right-left)/2 //防止整数溢出
		temp := find(matrix, mid)

		if temp >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func find(matrix [][]int, mid int) int {
	sum := 0
	n := len(matrix)
	i, j := n-1, 0
	for i >= 0 && j < n {
		if matrix[i][j] <= mid {
			sum += i + 1 //求出这一列的总数
			j++
		} else {
			i--
		}
	}

	return sum
}

func main() {
	matrix := [][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15}}
	k := 8
	fmt.Println(kthSmallest(matrix, k))
}
