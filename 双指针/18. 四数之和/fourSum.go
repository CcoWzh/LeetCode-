package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0)
	n := len(nums)
	if n < 4 {
		return result
	}
	//先排序
	sort.Ints(nums)
	//四个指针
	for i := 0; i < n; {
		//i,j也是指针
		for j := i + 1; j < n; {
			left, right := j+1, n-1
			//移动左右指针
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					//跳过重复的值
					cur3, cur4 := nums[left], nums[right]
					for left < n && nums[left] == cur3 {
						left++
					}
					for right >= 0 && nums[right] == cur4 {
						right--
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}

			cur2 := nums[j]
			for j < n && nums[j] == cur2 {
				j++
			}

		}
		//跳过重复的值
		cur1 := nums[i]
		for i < n && nums[i] == cur1 {
			i++
		}
	}

	return result
}

func main() {
	nums := []int{-3, -3, -3, -2, -1, -1, 0, 1, 1, 3, 4, 5, 5}
	target := 2
	fmt.Println(fourSum(nums, target))
}
