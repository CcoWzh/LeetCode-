package main

import "fmt"

func plusOne(digits []int) []int {
	mark := 1
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		sum := mark + digits[i]
		mark = sum / 10
		digits[i] = sum % 10
	}
	if mark == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {
	digits := []int{4, 3, 2, 1}
	fmt.Println(plusOne(digits))
}
