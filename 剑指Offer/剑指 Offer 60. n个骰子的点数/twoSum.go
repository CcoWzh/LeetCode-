package main

import "fmt"

//谁能想到是动态规划
//f(n) = 前一个数组的n-1~n-6的和
func twoSum(n int) []float64 {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 6*n+1)
	}

	//初始化，n=1时，点数和
	for i := 1; i < 7; i++ {
		dp[0][i] = 1
	}
	//fmt.Println(dp)
	count := 2
	for count <= n {
		cueMin, curMax := 1*count, 6*count
		for i := cueMin; i <= curMax; i++ {
			for j := 1; j <= 6; j++ {
				if i-j >= 1 {
					dp[count-1][i] += dp[count-2][i-j]
				}
			}
		}
		count++
	}
	fmt.Println(dp)

	cueMin, curMax := 1*n, 6*n
	res := make([]float64, curMax-cueMin+1)
	p := 6
	//6的n次方，这个都能出错？
	for i := 1; i < n; i++ {
		p *= 6
	}
	fmt.Println(p)

	for i := cueMin; i <= curMax; i++ {
		res[i-cueMin] = float64(dp[n-1][i]) / float64(p)
	}

	return res
}

func main() {
	n := 3
	fmt.Println(twoSum(n))
}
