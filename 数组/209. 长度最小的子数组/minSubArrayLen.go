package main

import "fmt"

func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	//求前缀和
	preSum := make([]int, n+1)
	preSum[0] = 0
	for i := 1; i < n+1; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	//求最小值
	result := preSum[n]
	//忘记加边界条件了
	if result < s {
		return 0
	}

	for i := 1; i < n+1; i++ {
		left, right := 0, i
		if preSum[right] < s {
			continue
		}
		for left < right {
			temp := preSum[right] - preSum[left]
			if temp < s {
				break
			} else {
				//fmt.Println(nums[left:right])
				if result > right-left {
					result = right - left
				}
				left++
			}
		}
	}

	return result
}

func main() {
	s := 16
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(s, nums))
}
