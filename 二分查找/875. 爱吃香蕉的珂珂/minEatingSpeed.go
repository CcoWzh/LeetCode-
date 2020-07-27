package main

import "fmt"

func minEatingSpeed(piles []int, H int) int {
	n := len(piles)
	left, right := 1, 0
	for i := 0; i < n; i++ {
		if right < piles[i] {
			right = piles[i]
		}
	}

	for left < right {
		mid := (left + right) / 2
		if isEating(piles, mid, H) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func isEating(piles []int, v, H int) bool {
	n, count := len(piles), 0
	for i := 0; i < n; i++ {
		//(piles[i]-1),是为了避免v=piles[i]的这种情况
		count += (piles[i]-1)/v + 1
	}
	//应该加=号，因为=H，可能不是最小值
	return count <= H
}

func main() {
	piles := []int{30, 11, 23, 4, 20}
	H := 6
	fmt.Println(minEatingSpeed(piles, H))
	fmt.Println(isEating(piles, 23, H))
}
