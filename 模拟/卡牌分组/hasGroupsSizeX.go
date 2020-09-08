package main

import "fmt"

//914. 卡牌分组
//这不就是求最大公约数吗
func hasGroupsSizeX(deck []int) bool {
	m := make(map[int]int)
	var result int
	for i, v := range deck {
		m[v]++
		if i == len(deck)-1 {
			result = m[v]
		}
	}

	for _, v := range m {
		if v < result {
			v, result = result, v
		}

		for v%result != 0 {
			v, result = result, v%result
		}
	}
	return result >= 2
}

func main() {
	deck := []int{1, 1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3}
	fmt.Println(hasGroupsSizeX(deck))
}
