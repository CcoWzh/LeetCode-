package main

import "fmt"

//这个操作是算法中常见的，作用是消除数字 n 的二进制表示中的最后一个 1。
func hammingWeight(num uint32) int {
	res := 0
	//每一次与，都相当于将最后一个1变为0
	//如：111 —> 110 —> 100 —> 000
	for num != 0 {
		//消除数字 n 的二进制表示中的最后一个 1
		num = num & (num - 1)
		res++
	}
	return res
}

func main() {
	fmt.Println(hammingWeight(8))
}
