package main

import "fmt"

func singleNumber(nums []int) int {
	bitSum := make([]int, 32)
	//将每个数字的每一位都相加，这样，出现3次的数字，其位数之和一定是3的倍数，即
	//如果把所有出现3次的数字的二进制表示的每一位都分别加起来，那么每一位的和都能被3整除
	//不能被3整除的部分，就说明，是只出现一次
	for i := 0; i < len(nums); i++ {
		mark := 1
		for j := 31; j >= 0; j-- {
			if nums[i]&mark == mark {
				bitSum[j]++
			}
			mark = mark << 1
		}
	}
	//计算只出现一次的数字
	result := 0
	//二进制转十进制，其实就是ai*2^(i-1)+......+
	for i := 0; i < len(bitSum); i++ {
		result = result << 1
		result += bitSum[i] % 3
	}

	return result
}

func main() {
	nums := []int{9, 6, 7, 9, 7, 9, 7}
	fmt.Println(singleNumber(nums))
	fmt.Println(twoToten([]int{1, 1, 0}))
}

func twoToten(bit []int) int {
	result := 0
	for i := 0; i < len(bit); i++ {
		result = result << 1
		result += bit[i]
	}
	return result
}
