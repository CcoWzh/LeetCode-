package main

import "fmt"

func intToRoman(num int) string {
	digits := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	index := 0
	result := ""
	for num > 0 {
		if num < digits[index] {
			index++
		} else {
			num -= digits[index]
			result += symbol[index]
		}
	}

	return result
}

func main() {
	num := 58
	fmt.Println(intToRoman(num))
}
