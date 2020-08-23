package main

import "fmt"

/** 二进制和位运算
1 << uint(len(nums)) ,即：2的n次方
i >> uint(j) & 1  即：i的二进制表示，用1来逐个遍历,
如101表示选择数组中0号和2号位置
 */
func subsets(nums []int,k int) [][]int {
	result := make([][]int, 0)

	for i := 0; i < (1 << uint(len(nums))); i++ {
		temp := make([]int, 0)
		for j := 0; j < len(nums); j++ {
			if (i >> uint(j) & 1) == 1 {
				temp = append(temp, nums[j])
			}
		}
		result = append(result, temp)
	}
	//fmt.Println(result)

	count :=0
	for i:=0;i< len(result);i++{
		if isRight(result[i],k){
			count++
			fmt.Println(result[i])
		}
	}
	fmt.Println(count)

	return result
}

func isRight(nums []int,k int) bool {
	n := len(nums)
	if n <=1{
		return true
	}

	for i:=0;i<n;i++{
		for j:=i+1;j<n;j++{
			if (nums[i]+nums[j]) %k ==0{
				return false
			}
		}
	}
	return true
}

func main() {
	nums := []int{4, 2, 3,1}
	fmt.Println(subsets(nums,3))
}

//func subsets(nums []int) [][]int {
//	// 保存最终结果
//	result := make([][]int, 0)
//	// 保存中间结果
//	list := make([]int, 0)
//	backtrack(nums, 0, list, &result)
//	return result
//}

// nums 给定的集合
// pos 下次添加到集合中的元素位置索引
// list 临时结果集合(每次需要复制保存)
// result 最终结果
func backtrack(nums []int, pos int, list []int, result *[][]int) {
	// 把临时结果复制出来保存到最终结果
	ans := make([]int, len(list))
	copy(ans, list)
	*result = append(*result, ans)
	// 选择、处理结果、再撤销选择
	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		backtrack(nums, i+1, list, result)
		list = list[0 : len(list)-1]
	}
}
