package main

import "fmt"

func maxDepthAfterSplit(seq string) []int {
	result := make([]int, len(seq))
	temp := make([]string, 0)
	for i := 0; i < len(seq); i++ {
		if string(seq[i]) == "(" {
			temp = append(temp, string(seq[i])) //入栈
			result[i] = len(temp) % 2
		} else {
			result[i] = len(temp) % 2
			temp = temp[:len(temp)-1] //出栈
		}
	}
	return result
}

func main() {
	seq := "(()(())())"
	fmt.Println(maxDepthAfterSplit(seq))

	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(nums[:len(nums)-1])
	fmt.Println(nums[len(nums)-1:])
}
