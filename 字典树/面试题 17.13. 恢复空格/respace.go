package main

import "fmt"

func respace(dictionary []string, sentence string) int {
	n := len(sentence)
	dp := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		dp[i] = n + 1
	}
	wordMap := make(map[string]bool)
	for i := 0; i < len(dictionary); i++ {
		wordMap[dictionary[i]] = true
	}

	for i := 1; i < n+1; i++ {
		dp[i] = dp[i-1] + 1
		for j := 0; j < i; j++ {
			if wordMap[sentence[j:i]] {
				dp[i] = min(dp[i], dp[j])
			}
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a > b {
		a = b
	}
	return a
}

func main() {
	dictionary := []string{"looked", "just", "like", "her", "brother"}
	sentence := "jesslookedjustliketimherbrother"
	fmt.Println(respace(dictionary, sentence))
}
