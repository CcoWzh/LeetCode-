package main

import (
	"fmt"
	"sort"
)

//由于冷却时间的存在，出现次数较多的那些任务如果不尽早安排，将会导致大量空闲时间的出现，
//因此贪心地将出现次数较多的任务安排在前面是合理的。
//func leastInterval(tasks []byte, n int) int {
//	count := make([]int, 26)
//	for i := 0; i < len(tasks); i++ {
//		count[tasks[i]-'A']++
//	}
//	//升序排序
//	sort.Ints(count)
//	time := 0
//	for count[25] > 0 {
//		i := 0
//		for i < n+1 {
//			if count[25] == 0 {
//				break
//			}
//
//			if i < 26 && count[25-i] > 0 {
//				count[25-i]--
//			}
//			time++
//			i++
//		}
//		//对当前计数从新排序
//		sort.Ints(count)
//	}
//	fmt.Println(time)
//	return time
//}

func main() {
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n := 2
	leastInterval(tasks, n)
}

func leastInterval(tasks []byte, n int) int {
	count := make([]int, 26)
	for _, v := range tasks {
		count[v-'A']++
	}
	sort.Ints(count)

	max := count[len(count)-1]
	blank := (max - 1) * n

	fmt.Println(max,blank)
	for i := 0; i < len(count)-1; i++ {
		if blank <= 0 {
			return len(tasks)
		}
		if count[i] == max {
			blank -= max - 1
		} else {
			blank -= count[i]
		}
	}
	return blank + len(tasks)
}