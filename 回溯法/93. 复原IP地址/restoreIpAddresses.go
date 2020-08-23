package main

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	result := make([]string, 0)
	//如果长度不够，不搜索
	if len(s) > 12 || len(s) < 4 {
		return result
	}
	backTrack(s, "", 0, &result)
	return result
}

//s 可做的选择列表，path当前做的选择，count 计数
func backTrack(s, path string, count int, result *[]string) {
	if count == 4 && len(s) == 0 {
		*result = append(*result, path[:len(path)-1])
	}

	temp := path
	for i := 1; i <= len(s); i++ {
		part := s[:i]
		if !isIP(part) {
			break
		}
		path = path + s[:i] + "." //做选择
		backTrack(s[i:], path, count+1, result)
		path = temp //退回选择
	}

}

//判断范围的
func isIP(part string) bool {
	//判断part不能是001等，一定要是合法的片段
	////大于 1 位的时候，不能以 0 开头
	if part[0] == '0' && len(part) != 1 {
		return false
	}

	ip, _ := strconv.Atoi(part)
	if ip < 0 || ip > 255 {
		return false
	}
	return true
}

func main() {
	s := "111111111111"
	fmt.Println(restoreIpAddresses(s))
}
