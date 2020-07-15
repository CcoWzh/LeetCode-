package main

import "fmt"

func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 {
		return nil
	}
	if shorter == longer {
		return []int{shorter * k}
	}
	result := make([]int, 0)
	//直接数学推导即可，没必要想的太复杂
	for i := 0; i <= k; i++ {
		f := shorter*k + (longer-shorter)*i
		result = append(result, f)
	}

	return result
}

func main() {
	shorter := 1
	longer := 1
	k := 5
	fmt.Println(divingBoard(shorter, longer, k))
}
