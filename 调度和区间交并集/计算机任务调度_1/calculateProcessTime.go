package main

import "fmt"

/**输入：n,m
		t[m]
 */
func calculateProcessTime(time []int, n, m int) []int {
	if time == nil {
		return nil
	}

	spendTime := make([]int, m)

	for i := 0; i < n; i++ {
		//给一个初始值，最好是很大的数
		minTime := spendTime[0] + time[0]
		minIndex := 0
		//贪心算法，比较每一个服务器，找到当前时间+执行一个任务后的时间，所需最少的那台服务器
		for j := 0; j < m; j++ {
			curTime := spendTime[j] + time[j]
			if curTime < minTime {
				minTime = curTime
				minIndex = j
			}
		}
		//更新所需时间
		spendTime[minIndex] += time[minIndex]
	}
	return spendTime
}

func main() {
	time := []int{7, 10}
	proTime := calculateProcessTime(time, 6, len(time))
	if proTime == nil {
		fmt.Println("分配失败")
		return
	}
	fmt.Println(proTime)
}
