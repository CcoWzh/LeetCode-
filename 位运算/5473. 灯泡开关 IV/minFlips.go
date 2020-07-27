package main

import "fmt"

//func minFlips(target string) int {
//	n := len(target)
//	t := make([]int, n+1)
//	t[0] = -1
//	for i := 1; i < n+1; i++ {
//		if target[i-1]-'0' == 1 {
//			t[i] = 1
//		} else {
//			t[i] = -1
//		}
//	}
//	//fmt.Println(t)
//	result := 0
//	for true {
//		result++
//		o := 0
//		count, i := 0, n
//		for i > 0 {
//			if t[i] == 1 && t[i-1] == -1 {
//				count++
//				break
//			}
//			if t[i] == -1 {
//				o++
//			}
//			count++
//			i--
//		}
//		if o == n {
//			break
//		}
//		for j := 0; j < count; j++ {
//			t[i] = 0 - t[i]
//			i++
//		}
//		//fmt.Println(t)
//	}
//
//	return result-1
//}

func main() {
	target := "001011101"
	fmt.Println(minFlips(target))
}

func minFlips(target string) int {
	target = "0" + target
	res := 0

	for i := 1; i < len(target); i++ {
		if target[i] != target[i-1] {
			res++
		}
	}
	return res
}
