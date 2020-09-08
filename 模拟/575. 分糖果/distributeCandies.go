package main

import "fmt"

func distributeCandies(candies []int) int {
	n := len(candies)
	m := make(map[int]int)

	count := 0
	for i := 0; i < n; i++ {
		if _, ok := m[candies[i]]; !ok {
			count++
		}
		m[candies[i]]++
	}

	if count <= n/2 {
		return count
	} else {
		return n / 2
	}
}

func main() {
	candies := []int{1, 1, 2, 3}
	fmt.Println(distributeCandies(candies))
}
