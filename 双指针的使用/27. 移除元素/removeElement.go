package main

import "fmt"

func removeElement(nums []int, val int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	count := 0
	for i := 0; i < n; i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		}
	}
	fmt.Println(nums[:count])
	return count
}

func main() {
	nums := []int{0,1,2,2,3,0,4,2}
	val := 2
	fmt.Println(removeElement(nums, val))
}
