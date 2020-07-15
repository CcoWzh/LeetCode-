package main

import "fmt"

func findContinuousSequence(target int) [][]int {
	result := make([][]int, 0)

	for i := 1; i < target/2+1; i++ {
		temp := make([]int, 0)
		s := 0
		for j := i; j < target; j++ {
			s = s + j
			temp = append(temp, j)
			if s == target {
				result = append(result, temp)
			}
			if s >target {
				break
			}
		}
	}

	return result
}

func main() {
	fmt.Println(findContinuousSequence(9))
}
