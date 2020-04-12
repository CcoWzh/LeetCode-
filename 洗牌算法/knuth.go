package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
Knuth-Shuffle，即 Knuth 洗牌算法。不就是蓄水池抽样算法嘛
对于生成的排列，每一个元素都能等概率的出现在每一个位置。
或者反过来，每一个位置都能等概率的放置每个元素。
 */
func knuth(nums []int) []int {

	for i := len(nums) - 1; i >= 0; i-- {
		rand.Seed(time.Now().Unix() + int64(i))
		j := rand.Int() % (i + 1)
		nums[i], nums[j] = nums[j], nums[i]
		//fmt.Println(nums, i, j)
	}
	//fmt.Println("======================")
	fmt.Println(nums)
	return nums
}

/**
rand.Intn(n int)
返回[0,n)之间的随机数
 */
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	knuth(nums)
}
