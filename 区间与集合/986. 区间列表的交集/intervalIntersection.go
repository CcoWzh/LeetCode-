package main

import "fmt"

func intervalIntersection(A [][]int, B [][]int) [][]int {
	result := make([][]int, 0, len(A)+len(B))
	i, j := 0, 0 //双指针
	for i < len(A) && j < len(B) {
		a1, a2 := A[i][0], A[i][1]
		b1, b2 := B[j][0], B[j][1]
		if b1 <= a2 && a1 <= b2 {
			//最重要的是，可以看到
			//{max(a1, b1), min(a2, b2)}这个关系
			result = append(result, []int{max(a1, b1), min(a2, b2)})
		}
		if a2 > b2 {
			j++
		} else {
			i++
		}
	}
	return result
}

func max(a, b int) int {
	if a < b {
		a = b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		a = b
	}
	return a
}

//A,B已经排序
func main() {
	A := [][]int{
		{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	B := [][]int{
		{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	fmt.Println(intervalIntersection(A, B))
}
