package main

import (
	"fmt"
	"sort"
)

func twoCitySchedCost(costs [][]int) int {
	n := len(costs)
	if n%2 == 1 || n < 1 {
		return 0
	}

	sort.Slice(costs, func(i, j int) bool {
		return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1]
	})
	sum := 0
	for i := 0; i < n/2; i++ {
		sum += costs[i][0] + costs[n-i-1][1]
	}

	return sum
}

func main() {
	costs := [][]int{{10, 20}, {30, 200}, {400, 50}, {30, 20}}
	fmt.Println(twoCitySchedCost(costs))
}
