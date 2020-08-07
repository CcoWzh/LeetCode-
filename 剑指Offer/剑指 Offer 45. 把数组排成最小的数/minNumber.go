package main

import (
	"fmt"
	"strconv"
	"strings"
)

func minNumber(nums []int) string {
	n := len(nums)
	//fmt.Println(nums)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			a := strconv.Itoa(nums[j])
			b := strconv.Itoa(nums[j+1])
			if strings.Compare(a+b, b+a) >= 0 { //a+b >= b+a ==> b<a
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	//fmt.Println(nums)
	var result string
	for i := 0; i < n; i++ {
		result += strconv.Itoa(nums[i])
	}

	return result
}

func main() {
	nums := []int{3, 30, 0, 5, 0}
	fmt.Println(minNumber(nums))
}
