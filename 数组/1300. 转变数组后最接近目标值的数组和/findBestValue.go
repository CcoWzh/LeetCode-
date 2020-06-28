package main

import (
	"fmt"
	"sort"
)

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	n := len(arr)

	preSum := make([]int, n+1)
	preSum[0] = 0
	for i := 1; i <= n; i++ {
		preSum[i] = arr[i-1] + preSum[i-1]
	}
	fmt.Println(preSum)

	res := target
	ans := 0
	for i := 1; i < arr[n-1]; i++ {
		index := sort.SearchInts(arr, i)
		sum := i*(n-index) + preSum[index]
		temp := abs(target - sum)

		if temp < res {
			res = temp
			ans = i
		}
	}

	return ans
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func main() {
	arr := []int{3, 2, 5}
	target := 10

	//arr := []int{60864, 25176, 27249, 21296, 20204}
	//target := 56803
	fmt.Println(findBestValue(arr, target))
}
