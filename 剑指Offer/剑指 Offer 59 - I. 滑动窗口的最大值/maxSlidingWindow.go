package main

import "fmt"

//你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}

	indexWin := make([]int, 0)
	for i := 0; i < k; i++ {
		//当新来的元素，比队尾元素大时，队尾元素出队列
		for len(indexWin) != 0 && nums[indexWin[len(indexWin)-1]] <= nums[i] {
			indexWin = indexWin[:len(indexWin)-1] //尾部，出队列
		}
		//添加新元素
		indexWin = append(indexWin, i)
	}

	maxWindows := make([]int, 0)
	for i := k; i < n; i++ {
		index := indexWin[0]
		maxWindows = append(maxWindows, nums[index])

		for len(indexWin) != 0 && nums[indexWin[len(indexWin)-1]] <= nums[i] {
			indexWin = indexWin[:len(indexWin)-1] //尾部，出队列
		}
		//当队首元素的下标，不在当前窗口中，则将队首元素出队列
		if i-k >= index && len(indexWin) != 0 {
			indexWin = indexWin[1:] //首部，出队列
		}
		indexWin = append(indexWin, i)
	}
	maxWindows = append(maxWindows, nums[indexWin[0]])

	return maxWindows
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(maxSlidingWindow(nums, k))
}
