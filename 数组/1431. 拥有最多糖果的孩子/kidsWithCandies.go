package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)
	max := 0
	for i := 0; i < n; i++ {
		if max < candies[i] {
			max = candies[i]
		}
	}

	result := make([]bool, n)

	for i := 0; i < n; i++ {
		if candies[i]+extraCandies >= max {
			result[i] = true
		} else {
			result[i] = false
		}
	}

	return result
}

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	fmt.Println(kidsWithCandies(candies, extraCandies))
}
