package main

import (
	"fmt"
	"sort"
)

var result [][]int

//39. 组合总和
func combinationSum(candidates []int, target int) [][]int {
	result = make([][]int, 0)
	sort.Ints(candidates)

	backtrack([]int{}, candidates, target, 0, 0)

	return result
}

func backtrack(track []int, candidates []int, target, cur, index int) {
	if cur == target {
		fmt.Println(track)
		result = append(result, track)
	}

	for i := index; i < len(candidates); i++ {
		temp := cur + candidates[i]
		if temp > target {
			continue
		}
		buf := make([]int, len(track))
		copy(buf, track) //----都得先保存当前路径
		buf = append(buf, candidates[i])
		backtrack(buf, candidates, target, temp, i)
	}

}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))
}
