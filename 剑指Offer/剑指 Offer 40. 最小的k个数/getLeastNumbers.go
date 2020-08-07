package main

import "fmt"

//快排和堆排
func getLeastNumbers(arr []int, k int) []int {
	n := len(arr)
	if n <= 0 || k <= 0 || k > n {
		return nil
	}

	nums := arr[:k]
	for i := len(nums)/2 - 1; i >= 0; i-- {
		sink(nums, i)
	}
	fmt.Println(nums)

	for i := k; i < n; i++ {
		if nums[0] <= arr[i] {
			//堆顶是最大值，最大值比arr[i]小，则说明当前arr[i]不在答案的序列中
			continue
		} else {
			nums[0] = arr[i]
			sink(nums, 0)
		}
	}

	return nums
}

//sink 下沉
func sink(nums []int, i int) {
	n := len(nums)
	for k := 2*i + 1; k < n; k = 2*k + 1 {
		if k+1 < n && nums[k] < nums[k+1] {
			k++
		}
		if nums[k] < nums[i] {
			break
		}
		//交换
		nums[i], nums[k] = nums[k], nums[i]
		i = k
	}
}

func main() {
	arr := []int{3, 1, 1, 9, 6}
	k := 2
	fmt.Println(getLeastNumbers(arr, k))
}
