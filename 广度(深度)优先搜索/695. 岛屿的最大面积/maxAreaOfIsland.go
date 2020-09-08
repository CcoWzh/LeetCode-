package main

import "fmt"

//695. 695. 岛屿的最大面积
func maxAreaOfIsland(grid [][]int) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			ans = max(dfs(grid, i, j), ans)
		}
	}
	return ans
}

func dfs(grid [][]int, i, j int) int {
	//截止条件
	if i < 0 || j < 0 || i == len(grid) || j == len(grid[0]) || grid[i][j] == 0 {
		return 0
	}
	grid[i][j] = 0 //搜索过的值，置为0
	ans := 1
	di := [4]int{0, 0, 1, -1}
	dj := [4]int{1, -1, 0, 0}
	for index := 0; index < 4; index++ {
		next_i, next_j := i+di[index], j+dj[index]
		ans += dfs(grid, next_i, next_j) //递归遍历
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		a = b
	}
	return a
}

func main() {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}

	//fmt.Println("i is : ", len(grid))
	//fmt.Println("j is : ", len(grid[0]))

	fmt.Println("岛屿的最大面积是:", maxAreaOfIsland(grid))
}

//func maxAreaOfIsland(grid [][]int) int {
//	ans := 0
//	for i := 0; i < len(grid); i++ {
//		for j := 0; j < len(grid[0]); j++ {
//			cur := 0
//			stack_i := make([]int, 0)
//			stack_j := make([]int, 0)
//			stack_i = append(stack_i, i)
//			stack_j = append(stack_j, j)
//
//			for len(stack_i) != 0 {
//				cur_i, cur_j := stack_i[0], stack_j[0]
//				stack_i = stack_i[1:]
//				stack_j = stack_j[1:]
//				if cur_i < 0 || cur_j < 0 || cur_i == len(grid) || cur_j == len(grid[0]) || grid[cur_i][cur_j] == 0 {
//					continue
//				}
//				cur++
//				grid[cur_i][cur_j] = 0
//				di := [4]int{0, 0, 1, -1}
//				dj := [4]int{1, -1, 0, 0}
//				for index := 0; index < 4; index++ {
//					next_i, next_j := cur_i+di[index], cur_j+dj[index]
//					stack_i = append(stack_i, next_i)
//					stack_j = append(stack_j, next_j)
//				}
//			}
//			ans = max(ans, cur)
//		}
//	}
//	return ans
//}
