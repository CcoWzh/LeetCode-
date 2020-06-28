package main

import "fmt"

func main() {
	A := []int{4, 5, 0, -2, -3, 1}
	K := 5
	subarraysDivByK(A, K)
}

//前缀和+哈希表+mod运算
func subarraysDivByK(A []int, K int) int {
	recode := make(map[int]int)
	recode[0] = 1

	result, modSum := 0, 0
	for _, v := range A {
		modSum += v
		temp := (modSum%K + K) % K
		//if s, ok := recode[temp]; ok {
		//	result += s
		//}
		result += recode[temp]
		recode[temp]++
	}

	fmt.Println(result)
	return result
}

/**
//本题提高了数据量的要求，这种两层循环的暴力法在本题不能通过，会超出时间限制
func subarraysDivByK(A []int, K int) int {
	n := len(A)
	preSum := make([]int, n+1)
	preSum[0] = 0
	for i := 0; i < n; i++ {
		preSum[i+1] += preSum[i] + A[i]
	}

	result := 0
	for i := 1; i < n+1; i++ {
		for j := 0; j < i; j++ {
			if (preSum[i]-preSum[j])%K == 0 {
				//fmt.Println(A[j:i])
				result++
			}
		}
	}
	//fmt.Println(result)
	return result
}
*/