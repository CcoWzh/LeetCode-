package main

import "fmt"

func fourSumCount(A []int, B []int, C []int, D []int) int {
	n := len(A)
	if n == 0 {
		return 0
	}
	mapAB := make(map[int]int)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			mapAB[A[i]+B[j]]++
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if mapAB[0-C[i]-D[j]] > 0 {
				count += mapAB[0-C[i]-D[j]]
			}
		}
	}

	return count
}

func main() {
	A := []int{1, 2}
	B := []int{-2, -1}
	C := []int{-1, 2}
	D := []int{0, 2}
	fmt.Println(fourSumCount(A, B, C, D))
}
