package main

import "fmt"

//必须在原数组上操作，不能拷贝额外的数组。
func moveZeroes(nums []int) {
	n := len(nums)
	fast, slow := 0, 0
	for fast < n {
		if nums[fast] == 0 { //当快指针的值是0时，直接跳过
			fast++
		} else {
			nums[slow] = nums[fast]
			fast++
			//慢指针指向的，最终是数组中不是0的个数
			//即，当前的非0元素的位置
			slow++
		}
	}

	for i := slow; i < n; i++ {
		nums[i] = 0
	}
	fmt.Println(nums)
}

func main() {
	nums := []int{1, 0, 3, 3, 3}
	moveZeroes(nums)
}
