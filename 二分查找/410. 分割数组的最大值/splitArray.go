package main

import "fmt"

func splitArray(nums []int, m int) int {
	n := len(nums)
	left, right := 0, 0
	//left : nums 中的最大值
	//right: nums 的数组和
	for i := 0; i < n; i++ {
		right += nums[i]
		if left < nums[i] {
			left = nums[i]
		}
	}

	for left < right {
		mid := (left + right) / 2
		//如果 cnt>m，说明划分的子数组多了，即我们找到的 mid 偏小，故 l=mid+1
		//否则，说明划分的子数组少了(标准值定大了)，即 mid 偏大(或者正好就是目标值)，故 r=mid
		if isSplit(nums, mid, m) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

//判断数组被x分割，会被分割成多少段
func isSplit(nums []int, x, m int) bool {
	sum, count := 0, 1
	//count从1开始取
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > x {
			//fmt.Println(sum)
			//当sum>x时，这个sum就已经不能取了，所以sum不是从0开始取的，说从nums[i]开始取
			sum = nums[i]
			count++
		}
	}
	return count <= m
}

func main() {
	nums := []int{7, 2, 5, 10, 8}
	m := 2
	fmt.Println(splitArray(nums, m))
}
