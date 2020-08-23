package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	if n == 0 {
		return -1
	}

	for i := 0; i < n; i++ {
		curSum := 0
		count := 0
		j := i
		for count < n && j < n {
			curSum = curSum + gas[j]
			if curSum < cost[j] {
				break
			}
			curSum = curSum - cost[j]
			j++
			if j >= n {
				j = 0
			}
			count++
		}

		if count == n {
			return i
		}
	}
	return -1
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
}
