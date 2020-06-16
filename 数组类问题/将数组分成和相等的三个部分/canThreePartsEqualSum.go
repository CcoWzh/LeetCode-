package main

import "fmt"

func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for _, v := range A {
		sum += v
	}
	if sum%3 != 0 {
		fmt.Println("sum is not 3")
		return false
	}
	temp, l := 0, 0
	for _, v := range A {
		temp += v
		if temp == sum/3 {
			l ++
			temp = 0
		}
		if l == 3 {
			return true
			break
		}
	}
	return false
}

func main() {
	//A := []int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}
	//A := []int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1}
	A := []int{1, -1, 1, -1}
	fmt.Println(canThreePartsEqualSum(A))
}
