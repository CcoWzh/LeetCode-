package main

import "fmt"

//139. 单词拆分
//动态规划
func wordBreak(s string, wordDict []string) bool {
	n, m := len(s), len(wordDict)
	if m == 0 {

	}
	//dp[i]表示以i结尾的字符串，是否可以完成拆分
	dp := make([]bool, n+1)
	wordMap := make(map[string]bool)
	for i := 0; i < m; i++ {
		wordMap[wordDict[i]] = true
	}
	dp[0] = true
	for i := 1; i < n+1; i++ {
		for j := 0; j < i; j++ {
			// 编程技巧
			// _leetcode,在字符串s前面，加一个空格，这样就不会产生s[j:i]的歧义
			// 下面的s[j:i]相当于s[j+1:i]
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
			}
		}
	}
	//fmt.Println(dp)
	return dp[n]
}

func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak(s, wordDict))
}
