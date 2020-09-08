package main

import "fmt"

func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	backtrack(nums, -1, 0, []int{}, &res)
	return res
}

// nums 给定的集合
// last 选择列表最后一个数字的下标
// pos 需要添加的数字的下标
// stack 临时结果集合(每次需要复制保存)
// result 最终结果
func backtrack(nums []int, last, pos int, stack []int, result *[][]int) {
	if len(nums) == pos {
		return
	}
	if (len(stack) == 0 || nums[pos] >= stack[len(stack)-1]) && isChack(nums, last, pos) {
		stack = append(stack, nums[pos]) //入栈
		if len(stack) >= 2 {
			ans := make([]int, len(stack))
			copy(ans, stack)
			*result = append(*result, ans)
		}

		backtrack(nums, pos, pos+1, stack, result)
		stack = stack[:len(stack)-1] //出栈
	}
	backtrack(nums, last, pos+1, stack, result)
}

func isChack(nums []int, last, pos int) bool {
	for i := last + 1; i < pos; i++ {
		if nums[i] == nums[pos] {
			return false
		}
	}
	return true
}

func main() {
	nums := []int{4, 6, 7, 7}
	fmt.Println(findSubsequences(nums))
}
