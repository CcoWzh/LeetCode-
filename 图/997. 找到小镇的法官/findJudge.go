package main

import "fmt"

func findJudge(N int, trust [][]int) int {
	//统计入度
	admission := make([]int, N+1)
	//统计出度
	out := make([]int, N+1)
	for i := 1; i < N+1; i++ {
		admission[i] = 0
		out[i] = 0
	}
	for i := 0; i < len(trust); i++ {
		from := trust[i][0]
		to := trust[i][1]
		admission[to] ++
		out[from] ++
	}
	count, index := 0, 0
	for i := 1; i < N+1; i++ {
		if admission[i] == N-1 {
			count++
			index = i
		}
	}

	fmt.Println(admission, out)

	if count != 1 {
		return -1
	}
	if out[index] != 0 {
		return -1
	}

	return index
}

func main() {
	N := 2
	trust := [][]int{}
	fmt.Println(findJudge(N, trust))
}

//更优雅的写法
//func findJudge(N int, trust [][]int) int {
//	Trust:=make([]bool,N+1)
//	Trusted:=make([]int,N+1)
//	Trust[0]=true
//	for _,info:=range trust {
//		Trust[info[0]]=true
//		Trusted[info[1]]++
//	}
//	for i:=1;i<=N;i++ {
//		if !Trust[i] && Trusted[i]==N-1 {
//			return i
//		}
//	}
//	return -1
//}
