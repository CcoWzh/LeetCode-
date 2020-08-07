package main

import "fmt"

/**
三指针法
 */
func nthUglyNumber(n int) int {
	index2, index3, index5 := 0, 0, 0
	dp := make([]int, 1)
	dp[0] = 1
	//1,2,3,4,5,6,8,9,10,12
	for i := 1; i < n; i++ {
		//设：当前求得的丑数为M
		//第一个乘以2后，大于M的数， 第一个乘以3后，大于M的数， 第一个乘以5后，大于M的数，
		//3个指针在dp数组上依次移动
		dp = append(dp, min(dp[index2]*2, dp[index3]*3, dp[index5]*5))
		if dp[i] == dp[index2]*2 {
			index2++
			//1*2,2*2,3*2,4*2,5*2,6*2,8*2......
		}
		if dp[i] == dp[index3]*3 {
			//1*3,2*3,3*3,4*3,5*3,6*3,8*3......
			index3++
		}
		if dp[i] == dp[index5]*5 {
			//也就是说，按照丑数的顺序，依次乘以5
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
