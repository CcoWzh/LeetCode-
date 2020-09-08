package main

import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	//排序
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] { //如果高度相等，则按照k值排序（升序）
			return people[i][1] < people[j][1]
		}
		//按照高度，降序
		return people[i][0] > people[j][0]
	})
	//按照k值插入到index=k的地方，index之后的往后移动
	result := make([][]int, 0)
	for _, info := range people {
		result = append(result, info)
		copy(result[info[1]+1:], result[info[1]:])
		result[info[1]] = info
	}

	return result
}

func main() {
	people := [][]int{
		{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	fmt.Println(reconstructQueue(people))
}
