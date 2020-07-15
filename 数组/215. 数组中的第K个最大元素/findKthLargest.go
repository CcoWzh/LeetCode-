package main

import "fmt"

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	fmt.Println(findKthLargest(nums, k))
}

//最小堆
func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return -1
	}
	//1.构建大顶堆
	for i := len(nums)/2 - 1; i >= 0; i-- {
		//从第一个非叶子结点从下至上，从右至左调整结构
		//第一个非叶子节点是：len(nums)/2 - 1
		sink(nums, i)
	}

	for i := 0; i < k-1; i++ {
		//删除堆顶元素
		n := len(nums)
		nums[0] = nums[n-1]
		nums = nums[:n-1]
		//将尾部的元素加入到顶部，然后将顶部元素下沉
		sink(nums, 0)
	}

	return nums[0]
}

//swin 上浮
/**
上浮，就是先找到这个元素的父节点，即(i - 1) / 2
和父节点比较，如果比父节点大，交换两个元素，继续比较

一般，构建最大堆的插入操作，会使用上浮
 */
func swin(nums []int, i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if nums[parent] < nums[i] {
			nums[parent], nums[i] = nums[i], nums[parent]
		} else {
			break
		}
		i = parent
	}
}

//sink 下沉
/**
下沉，就是从父节点出发，和自己的两个子节点比较。如果比最大的子节点还小，则和它交换，继续比较
直到超出界限，或者比左右节点都小。

一般，构建最大堆时，用于删除操作。将头部数据删除，复制尾部的数据到头部，下沉头部数据
 */
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
