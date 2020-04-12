package main

import "fmt"

/**
三指针法
 */
func nthUglyNumber(n int) int {
	index2, index3, index5 := 0, 0, 0
	dp := make([]int, 1)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp = append(dp, min(dp[index2]*2, dp[index3]*3, dp[index5]*5))
		if dp[i] == dp[index2]*2 {
			index2++
		}
		if dp[i] == dp[index3]*3 {
			index3++
		}
		if dp[i] == dp[index5]*5 {
			index5++
		}
	}
	fmt.Println(dp)
	return dp[len(dp)-1]
}

func min(a, b, c int) int {
	if a > b {
		a = b
	}
	if a > c {
		a = c
	}
	return a
}

func main() {
	fmt.Println(nthUglyNumber(10))
}
