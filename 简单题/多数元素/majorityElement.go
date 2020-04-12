package main

import "fmt"

//169. 多数元素
func majorityElement(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v] ++
		if m[v] > len(nums)/2 {
			return v
		}
	}

	return 0
}

func main() {
	nums := []int{3, 2, 3}
	fmt.Println(majorityElement(nums))
}
