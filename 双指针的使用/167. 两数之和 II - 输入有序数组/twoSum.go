package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	index := make([]int, 2)
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			break
		}
		if sum > target {
			right--
		} else if sum < target {
			left++
		}
	}
	index[0] = left + 1
	index[1] = right + 1
	return index
}

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(numbers, target))
}
