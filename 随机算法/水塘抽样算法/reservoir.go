package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sample(dataStream []int, m int) []int {
	n := len(dataStream)
	reservoir := make([]int, m)
	// init
	for i := 0; i < m; i++ {
		reservoir[i] = dataStream[i]
	}

	for i := m; i < n; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		// 随机获得一个[0, i]内的随机整数
		d := r.Intn(i + 1)
		// 如果随机整数落在[0, m-1]范围内，则替换蓄水池中的元素
		if d < m {
			reservoir[d] = dataStream[i]
		}
	}
	return reservoir
}

func main() {
	dataStream := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sample(dataStream, 5))
}
