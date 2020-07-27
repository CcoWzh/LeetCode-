package main

import "fmt"

/**
一个更好的解法，就是从后往前计数
 */
func romanToInt(s string) int {
	sum, n := 0, len(s)
	if n == 0 {
		return 0
	}
	wordMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i := 0; i < n; i++ {
		if i+1 < n && string(s[i]) == "I" && (string(s[i+1]) == "V" || string(s[i+1]) == "X") {
			sum += wordMap[string(s[i+1])] - 1
			i++
			continue
		}
		if i+1 < n && string(s[i]) == "X" && (string(s[i+1]) == "L" || string(s[i+1]) == "C") {
			sum += wordMap[string(s[i+1])] - 10
			i++
			continue
		}
		if i+1 < n && string(s[i]) == "C" && (string(s[i+1]) == "D" || string(s[i+1]) == "M") {
			sum += wordMap[string(s[i+1])] - 100
			i++
			continue
		}
		sum += wordMap[string(s[i])]
	}

	return sum
}

func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}
