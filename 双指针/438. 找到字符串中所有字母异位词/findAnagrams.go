package main

import "fmt"

//滑动窗口
func findAnagrams(s string, p string) []int {
	left, right := 0, 0
	window := make(map[uint8]int)
	need := make(map[uint8]int)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	valid := 0 //当valid有效时，更新答案
	res := make([]int, 0)
	for right < len(s) {
		cur := s[right]
		right++ //右移窗口
		//进行窗口内数据的一系列更新
		if _, ok := need[cur]; ok {
			window[cur]++
			if window[cur] == need[cur] {
				valid++
			}
		}
		//滑动窗口，窗口大小最大是len(p)
		if right-left == len(p) { //当子串是连续的时候
			if valid == len(need) {
				res = append(res, left)
			}
			d := s[left]
			//缩小窗口
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	return res
}

func main() {
	s := "abfcab"
	p := "abc"
	fmt.Println(findAnagrams(s, p))
}
