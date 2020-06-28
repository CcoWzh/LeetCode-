package main

import "fmt"

/**
  最关键的是知道拆分成：A[i] + i 和 A[j] - j
  max = A[i] + i
 */
func maxScoreSightseeingPair(A []int) int {
	result, maxScore := 0, 0
	for i := 0; i < len(A); i++ {
		result = max(result, maxScore+A[i]-i)
		maxScore = max(maxScore, A[i]+i)
	}
	return result
}

func max(a, b int) int {
	if a < b {
		a = b
	}
	return a
}

func main() {
	A := []int{8, 1, 5, 2, 6}
	fmt.Println(maxScoreSightseeingPair(A))
}

////状态：超出时间限制
//func maxScoreSightseeingPair(A []int) int {
//	max := 0
//	n := len(A)
//	for i := 0; i < n; i++ {
//		for j := i + 1; j < n; j++ {
//			temp := A[i] + A[j] + i - j
//			if max < temp {
//				max = temp
//			}
//		}
//	}
//
//	return max
//}

/* 更优美的写法
func maxScoreSightseeingPair(A []int) int {

	var max int
	var left int
	for i := 0; i < len(A); i++ {
		sum := A[i] - i + left
		if sum > max {
			max = sum
		}
		t := A[i] + i
		if left < t {
			left = t
		}
	}
	return max
}
*/
