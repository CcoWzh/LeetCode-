package main

import "fmt"

func permutation(s string) []string {
	n := len(s)
	s = bubbleSort(s)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {

		}
	}

	if n == 0 {
		return nil
	}
	result := make([]string, 0)
	backTrack(s, "", &result)
	return result
}

//s 选择列表, cur当前的选择
//count 当前进行到哪一步, all，s的总长度
func backTrack(s, cur string, result *[]string) {
	if len(s) == 0 {
		*result = append(*result, cur)
	}

	for i := 0; i < len(s); i++ {
		temp := s[i]
		cur = cur + string(temp) // 做选择
		backTrack(deleteS(s, i), cur, result)
		cur = cur[:len(cur)-1] //退回选择
		//需要跳过重复的字符
		j := i
		for j+1 < len(s) && s[j+1] == temp {
			i++
			j++
		}
	}

}

func bubbleSort(s string) string {
	arr := []byte(s)
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return string(arr)
}

//删除字符串中的第i个元素
//string 是深拷贝，拷贝的是数据本身，不会复制地址指针
func deleteS(s string, i int) string {
	n := len(s)
	if i >= n || i < 0 {
		return s
	}
	return s[:i] + s[i+1:]
}

func main() {
	s := "aabcc"
	fmt.Println(permutation(s))
}
