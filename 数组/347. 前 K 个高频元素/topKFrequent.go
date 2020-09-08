package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	n := len(nums)
	numMap := make(map[int]int)
	for i := 0; i < n; i++ {
		numMap[nums[i]]++
	}

	size := 0
	arr := make([][]int, 0) //扩展到二维上
	for i, v := range numMap {
		if size < k {
			arr = append(arr, []int{v, i})
			n := len(arr)
			swin(arr, n-1) //上浮
			size++
		} else {
			if v > arr[0][0] {
				//删除堆顶元素，将新入的元素加入堆顶，下沉
				arr[0] = []int{v, i}
				sink(arr, 0)
			}
		}
	}
	fmt.Println(arr)
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = arr[i][1]
	}
	return res
}

//sink 下沉
func sink(nums [][]int, i int) {
	n := len(nums)
	for k := 2*i + 1; k < n; k = 2*k + 1 {
		if k+1 < n && nums[k][0] > nums[k+1][0] { //比较两个左右孩子
			k++
		}
		if nums[k][0] > nums[i][0] {
			break
		}
		//交换
		nums[i], nums[k] = nums[k], nums[i]
		i = k
	}
}

//上浮
func swin(nums [][]int, i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if nums[parent][0] > nums[i][0] {
			nums[parent], nums[i] = nums[i], nums[parent]
		} else {
			break
		}
		i = parent
	}
}

func main() {
	nums := []int{1,1,1,1,1,2,2,3,4,4,4,4,5,5,5,5}
	k := 3
	fmt.Println(topKFrequent(nums, k))
}
