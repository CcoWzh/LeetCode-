package main

import "fmt"

func judgePoint24(nums []int) bool {
	numbers := make([]float64, len(nums))
	for i, v := range nums {
		numbers[i] = float64(v)
	}
	return dfs(numbers)
}

func dfs(nums []float64) bool {
	n := len(nums)
	if n == 1 {
		return abs(nums[0]-24) < 1e-6
	}
	//从numbers中取出两个数，把结果放入数组中
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j { //取不同的两个数
				//如果回溯的话，还要恢复现场，把数插回原位置，所以不如直接生成一个新数组
				temp := make([]float64, 0)
				for k := 0; k < n; k++ {
					if k == i || k == j {
						continue
					}
					temp = append(temp, nums[k])
				}
				res := calculate(nums[i], nums[j]) //获取两个数运算的结果集
				for _, v := range res {
					temp = append(temp, v)
					if dfs(temp) {
						return true
					}
					temp = temp[:len(temp)-1] //恢复现场
				}

			}
		}
	}
	return false
}

//计算两个数的运算结果
func calculate(a, b float64) []float64 {
	res := make([]float64, 0)
	res = append(res, a+b)
	res = append(res, a*b)
	res = append(res, a-b)
	res = append(res, b-a)

	if a > 0 {
		res = append(res, b/a)
	}
	if b > 0 {
		res = append(res, a/b)
	}
	return res
}

func abs(a float64) float64 {
	if a < 0 {
		a = -a
	}
	return a
}

func main() {
	nums := []int{3, 3, 8, 8}
	fmt.Println(judgePoint24(nums))
}
