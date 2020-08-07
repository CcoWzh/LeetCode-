package main

import (
	"fmt"
	"strconv"
)

func getPermutation(n int, k int) string {
	//对于数字n来说，是以每(n-1)!一组
	var factorial = []int{0, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880, 3628800}
	result := ""
	if k < 1 || k > factorial[n] {
		return result
	}

	nums := make([]string, n)
	//转换成字符串类型
	for i := 1; i <= n; i++ {
		nums[i-1] = strconv.Itoa(i)
	}
	//规律
	k--
	for k > 0 {
		index := k / factorial[n-1]
		result += nums[index]
		nums = deleteS(nums, index)
		k = k % factorial[n-1]
		n--
	}
	//当k==0，nums有剩余时，说明这个应该正序串成字符串
	for len(nums) != 0 {
		result += nums[0]
		nums = nums[1:]
	}

	return result
}

//这个要注意,是切片类型，会改变原先的s
func deleteS(s []string, i int) []string {
	if i >= len(s) {
		return nil
	}
	result := make([]string, len(s))
	copy(result, s)
	temp := result[i+1:]
	result = result[:i]
	result = append(result, temp...)
	return result
}

func main() {
	n, k := 4, 13
	fmt.Println(getPermutation(n, k))
}
