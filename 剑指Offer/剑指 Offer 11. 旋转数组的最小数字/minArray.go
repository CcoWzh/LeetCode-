package main

import "fmt"

func minArray(numbers []int) int {
	n := len(numbers)
	if n == 0 {
		return 0
	}
	min := numbers[0]

	for i := 0; i < n; i++ {
		if min > numbers[i] {
			min = numbers[i]
		}
	}
	return min
}

func main() {
	numbers := []int{3, 4, 5, 1, 2}
	fmt.Println(minArray(numbers))
}
