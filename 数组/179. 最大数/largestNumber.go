package main

import (
	"fmt"
	"strconv"
)

//重新定义数字的大小关系
func largestNumber(nums []int) string {
	n := len(nums)
	if n == 0 {
		return ""
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			a := strconv.Itoa(nums[j])   //6
			b := strconv.Itoa(nums[j+1]) //51
			if a+b > b+a { //a+b >= b+a ==> b<a  651>516 ==> 6>51
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	fmt.Println(nums)
	res := "" //这个需要去除前导0
	for i := n - 1; i >= 0; i-- {
		res += strconv.Itoa(nums[i])
	}

	for res[0] == '0' && len(res) > 1 {
		res = res[1:]
	}

	return res
}

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(largestNumber(nums))
}
