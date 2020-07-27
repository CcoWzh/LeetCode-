package main

import "fmt"

var wordMap = map[int]string{
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

func letterCombinations(digits string) []string {
	n := len(digits)
	if n == 0 {
		return []string{""}
	}
	result := make([]string, 0)
	backtrack(digits, "", 0, n, &result)
	return result
}

func backtrack(selectList, path string, curIndex, n int, result *[]string) {
	if curIndex == n {
		//s := path
		*result = append(*result, path)
		return
	} else {
		list := wordMap[int(selectList[curIndex]-'0')]

		for i := 0; i < len(list); i++ {
			path += string(list[i])
			backtrack(selectList, path, curIndex+1, n, result)
			path = path[:len(path)-1]
		}
	}
}

func main() {
	digits := "93"
	fmt.Println(letterCombinations(digits))
}
