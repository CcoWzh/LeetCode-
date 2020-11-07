package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	original []int //初始时的数组
	nums     []int //打乱后的数组
}

func Constructor(nums []int) Solution {
	temp := make([]int, len(nums))
	copy(temp, nums)
	return Solution{
		original: nums,
		nums:     temp}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	copy(this.nums, this.original)
	return this.nums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	len := len(this.nums)
	//洗牌算法
	for index, _ := range this.nums {
		//产生[index,len)之间的随机数
		swapIndex := random(index, len)
		this.nums[index], this.nums[swapIndex] = this.nums[swapIndex], this.nums[index]
	}
	return this.nums
}

// [min, max)
func random(min, max int) int {
	//[0,n)
	return rand.Intn(max-min) + min
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
func main() {
	c := Constructor([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(c.Shuffle())
	fmt.Println(c.Reset())
	fmt.Println(c.Shuffle())
}
