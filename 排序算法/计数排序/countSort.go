package main

import "fmt"

func countSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	// 1.找出数组中最大数和最小数
	var maxValue, minValue int
	minValue = nums[0] //不加这一句的话，minValue永远都是0
	for _, v := range nums {
		if v > maxValue {
			maxValue = v
		}
		if v < minValue {
			minValue = v
		}
	}

	bucketLen := maxValue - minValue + 1
	bucket := make([]int, bucketLen)
	// 桶中每个位置用来计数原数组中等于"当前桶索引+minValue"的数的个数
	for _, v := range nums {
		bucket[v-minValue]++
	}
	// 将桶里的元素放回原数组
	sortedIndex := 0
	for i := 0; i < bucketLen; i++ {
		for bucket[i] > 0 {
			//由于bucket[v-minValue]++的保存方式，所以取出来时要+minValue
			nums[sortedIndex] = i + minValue
			sortedIndex++
			bucket[i]--
		}
	}

	fmt.Println(nums)
	return nums
}

func main() {
	//nums := []int{8, 1, 4, 6, 2, 3, 5, 7}
	nums := []int{6, 65, 58, 95, 10, 57, 62, 13, 106, 78, 23, 85}
	countSort(nums)
}
