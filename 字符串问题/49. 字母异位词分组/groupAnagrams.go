package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	n := len(strs)
	if n == 0 {
		return nil
	}
	wordNums := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	wordMap := make(map[int][]string)
	//a~z:是1~26
	for i := 0; i < n; i++ {
		//计算字符串转换成数字后的和
		sum := 1
		for j := 0; j < len(strs[i]); j++ {
			num := strs[i][j] - 'a'
			sum *= wordNums[num]
		}
		wordMap[sum] = append(wordMap[sum], strs[i])
	}
	//fmt.Println(wordMap)
	result := make([][]string, 0)
	for _, v := range wordMap {
		result = append(result, v)
	}

	return result
}

func main() {
	strs := []string{"has", "rod", "tom", "hum", "him", "yon", "met", "dye"}
	fmt.Println(groupAnagrams(strs))
}
