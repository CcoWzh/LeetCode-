package main

import "fmt"

func sortByBits(arr []int) []int {
	n := len(arr)
	bit := make([]int, n)

	for i := 0; i < n; i++ {
		temp, count := arr[i], 0
		for temp != 0 {
			temp = temp & (temp - 1)
			count++
		}
		bit[i] = count
	}
	fmt.Println(bit)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if bit[j] < bit[i] {
				bit[i], bit[j] = bit[j], bit[i]
				arr[i], arr[j] = arr[j], arr[i]
			} else if bit[j] == bit[i] {
				if arr[i] > arr[j] {
					arr[i], arr[j] = arr[j], arr[i]
				}
			}
		}
	}

	return arr
}

func main() {
	arr := []int{2}
	fmt.Println(sortByBits(arr))
}
