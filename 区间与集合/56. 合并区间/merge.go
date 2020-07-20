package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	n := len(intervals) //无需判断是否为空
	result := make([][]int, 0)
	//排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	//合并区间，双指针
	for i := 0; i < n; {
		left, right := intervals[i][0], intervals[i][1]
		for i+1 < n && right >= intervals[i+1][0] {
			if right <= intervals[i+1][1] {
				right = intervals[i+1][1]
				i++
			} else {
				i++
			}
		}
		result = append(result, []int{left, right})
		i++
	}

	return result
}

func main() {
	//intervals := [][]int{
	//	{1, 3}, {15, 18}, {8, 10}, {2, 6}}
	intervals := [][]int{
		{1, 4}, {2, 3}, {4, 5}, {7, 9}, {6, 7}, {11, 15}}
	fmt.Println(merge(intervals))

}
