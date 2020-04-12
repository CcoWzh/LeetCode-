package main

import (
	"fmt"
)

/**
桶排序原理是：将数组分到有限数量的桶子里（分治法）
然后对每个桶子再分别排序，最后将各个桶中的数据有序的合并起来。
 */
func bucketSort(nums []int) []int {
	// 1.找出数组中最大数和最小数
	var maxValue, minValue int
	for _, v := range nums {
		if v > maxValue {
			maxValue = v
		}
		if v < minValue {
			minValue = v
		}
	}
	// 2.创建桶，桶的个数为 3
	bucketSize := 3
	buckets := make([][]int, bucketSize)
	// 3.把数据分配到桶中，桶中的数据是有序的
	// a.计算桶中数据的平均值，这样分组数据的时候会把数据放到对应的桶中
	space := float64(maxValue-minValue+1) / float64(bucketSize)
	for _, v := range nums {
		index := int(float64(v-minValue) / space)
		buckets[index] = append(buckets[index], v)
		////b.对每个桶进行排序，插入排序
		//for i := 1; i < len(buckets[index]); i++ {
		//	for j := i; j > 0 && buckets[index][j] < buckets[index][j-1]; j-- {
		//		buckets[index][j], buckets[index][j-1] = buckets[index][j-1], buckets[index][j]
		//	}
		//}
	}
	fmt.Println(buckets)
	// b.对每个桶进行排序，插入排序
	for _, bucketv := range buckets {
		insertionSort(bucketv)
	}

	fmt.Println(buckets)
	// 4.把桶中的数据重新组装起来
	nums = []int{}
	for _, bucketv := range buckets {
		nums = append(nums, bucketv...)
	}

	fmt.Println(nums)
	return nums

}

func insertionSort(numArr []int) {
	for i := 0; i < len(numArr); i++ {
		for j := 0; j < len(numArr)-1; j++ {
			if numArr[j] > numArr[j+1] {
				numArr[j], numArr[j+1] = numArr[j+1], numArr[j]
			}
		}
	}
}

func main() {
	//nums := []int{8, 1, 4, 6, 2, 3, 5, 7}
	nums := []int{65, 58, 95, 10, 57, 62, 13, 106, 78, 23, 85}
	bucketSort(nums)
}
