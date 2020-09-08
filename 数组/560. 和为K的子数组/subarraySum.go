package main

import "fmt"

func subarraySum(nums []int, k int) int {
	//前缀和
	n := len(nums)
	preSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	//前缀求和，preSum[j]-preSum[i] == k
	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			//fmt.Println(nums[i:j], preSum[j], preSum[i])
			if preSum[j]-preSum[i] == k {
				count++
			}
		}
	}
	return count
}

func main() {
	nums := []int{-1,-1,1}
	k := 0
	fmt.Println(subarraySum(nums, k))
}
