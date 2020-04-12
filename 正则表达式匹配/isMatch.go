package main

import "fmt"

/**
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。
动态规划
 */
func isMatch(s string, p string) bool {
	if len(s) != 0 && len(p) == 0 {
		return false
	}

	m, n := len(s), len(p)

	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true

	//"" 和p的匹配关系初始化，a*a*a*a*a*这种能够匹配空串，其他的是都是false。
	//  奇数位不管什么字符都是false，偶数位为* 时则: dp[0][i] = dp[0][i - 2]
	for i := 2; i <= n; i += 2 {
		if string(p[i-1]) == "*" {
			dp[0][i] = dp[0][i-2]
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			sc := string(s[i-1])
			pc := string(p[j-1])
			if sc == pc || pc == "." {
				dp[i][j] = dp[i-1][j-1]
			} else if pc == "*" {
				if dp[i][j-2] {
					dp[i][j] = true;
				} else if sc == string(p[j-2]) || string(p[j-2]) == "." {
					dp[i][j] = dp[i-1][j]
				}
			}
		}
	}

	return dp[m][n]
}

func main() {
	s := ""
	p := ""

	fmt.Println(isMatch(s, p))

	//s := "mississippi"
	//p := "mis*is*p*."
	//fmt.Println(isMatch(s, p))
}
