package main

import "fmt"

//1160. 拼写单词
func countCharacters(words []string, chars string) int {
	if len(words) == 0 || len(chars) == 0 {
		return 0
	}
	count := 0
	m := make(map[string]int)
	for _, v := range chars {
		m[string(v)] ++
	}

	for i := 0; i < len(words); i++ {
		c := make(map[string]int)
		temp := 0
		for j := 0; j < len(words[i]); j++ {
			c[string(words[i][j])]++
			if v, ok := m[string(words[i][j])]; ok && v >= c[string(words[i][j])] {
				temp++
			} else {
				break
			}
		}
		if temp == len(words[i]) {
			count += len(words[i])
		}
	}
	return count
}

func main() {
	words := []string{"hello","world","leetcode"}
	chars := "welldonehoneyr"
	fmt.Println(countCharacters(words, chars))
}
