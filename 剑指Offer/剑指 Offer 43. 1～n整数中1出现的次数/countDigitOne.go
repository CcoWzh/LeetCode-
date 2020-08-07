package main

import "fmt"

//教程：
//https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof/solution/gui-lu-ti-ji-lu-shu-duan-jie-guo-sheng-liao-bu-sha/
func countDigitOne(n int) int {
	record := []int{0, 1, 20, 300, 4000, 50000, 600000, 7000000, 80000000, 900000000}

	bitPow, result, pow10, tmpn := 0, 0, 1, n
	for tmpn != 0 {
		num := tmpn % 10
		tmpn = tmpn / 10
		result += record[bitPow] * num //+2*20
		if num > 1 {
			result += pow10 // +100,记录当前最高位有多少个1
		} else if num == 1 {  //排除0的情况
			result += n%pow10 + 1 //161----62个1
		}
		bitPow++
		pow10 *= 10 //10的n次方
	}
	return result
}

func main() {
	fmt.Println(countDigitOne(10))
}
