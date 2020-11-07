package main

import "fmt"

//945. 使数组唯一的最小增量
func minIncrementForUnique(A []int) int {
	if len(A) == 0 {
		return 0
	}
	result, count := 0, 0
	m := make(map[int]int)
	min, max := A[0], A[0]
	for _, v := range A {
		if _, ok := m[v]; ok {
			count++
		}
		m[v]++
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}

	temp := make([]int, 0)
	//列举出没有出现过的数
	for i := min; i <= max+count*2; i++ {
		if _, ok := m[i]; !ok {
			temp = append(temp, i)
		}
	}

	//fmt.Println(temp)
	for i, v := range m {
		if v == 1 {
			continue
		}

		for j := 0; j < len(temp) && v != 1; {
			if temp[j] > i {
				result += temp[j] - i
				v--
				temp = append(temp[0:j], temp[j+1:]...)
			} else {
				j++
			}
		}
		//fmt.Println(temp)
	}

	return result
}

func main() {
	A := []int{3, 2, 1, 2, 1, 7}
	//3,2,1,2,1,7
	fmt.Println(minIncrementForUnique(A))
}
