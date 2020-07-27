package main

import "fmt"

func restoreString(s string, indices []int) string {
	n := len(s)
	result := make([]string, n)

	for i := 0; i < n; i++ {
		index := indices[i]
		result[index] = string(s[i])
	}
	fmt.Println(result)

	ss := ""
	for i := 0; i < n; i++ {
		ss += result[i]
	}
	return ss
}

//func restoreString(s string, indices []int) string {
//	r, n := []byte(s), len(indices)
//	for i := 0; i < n/2; i++ {
//		index := indices[i]
//		r[i], r[index] = r[index], r[i]
//	}
//
//	return string(r)
//}

func main() {
	//s := "codeleet"
	//indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	//fmt.Println(restoreString(s, indices))

	s := "aaiougrt"
	indices := []int{4, 0, 2, 6, 7, 3, 1, 5}
	fmt.Println(restoreString(s, indices))
}
