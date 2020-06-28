package main

import (
	"fmt"
	"math"
	"sort"
)

//16. 最接近的三数之和
//双指针法，和15.三数之和的解法相似
func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)

	//result := nums[0] + nums[1] + nums[2]
	//dis这个变量，好像也可以去掉的吧
	dis, result := math.MaxInt32, math.MaxInt32
	for i := 0; i < n; i++ {
		//双指针
		left, right := i+1, n-1
		for left < right {
			_sum := nums[i] + nums[left] + nums[right]
			//最接近target
			if abs(target-_sum) < dis {
				dis = abs(target - _sum) //这个忘了加绝对值了
				result = _sum
			}
			//移动指针
			if _sum == target {
				return _sum
			} else if _sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

//绝对值
func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func main() {
	nums := []int{-1, 0, 1, 1, 55}
	target := 3
	fmt.Println(threeSumClosest(nums, target))
}
