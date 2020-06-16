package main

import "fmt"

/**
切片，是左开右闭的，也就是说，nums[i:j] --> i<=x<j
 */
func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(nums[0])
	fmt.Println(nums[:3])
	fmt.Println(nums[1:3])
	fmt.Println(len(nums)-1)
	fmt.Println(nums[3:7])
}
