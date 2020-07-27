package main

import "fmt"

func pathInZigZagTree(label int) []int {
	count, my := 0, label
	for my > 1 {
		my = my >> 1
		count++
	}
	//构建树
	tree := make([]int, 1<<uint(count+1)-1)
	i, j, k := 0, 0, 0
	for j < count+1 {
		if j%2 == 0 {
			tree[i] = i + 1
		} else {
			tree[i] = 1<<uint(j) + 1<<uint(j+1) - 2 - i
		}
		k++
		i++
		if k > 1<<uint(j)-1 {
			j++
			k = 0
		}
	}
	//寻找label的下标
	index := label
	if count%2 == 1 {
		left, right := 1<<uint(count), 1<<uint(count+1)
		a := right - label
		b := left - 2
		index = b + a
	} else {
		index = index - 1
	}
	//fmt.Println(tree, index)
	//大堆顶，上浮
	path := make([]int, count+1)
	i = count
	for i >= 0 {
		path[i] = tree[index]
		index = (index - 1) / 2
		i--
	}

	return path
}

func main() {
	label := 10
	fmt.Println(pathInZigZagTree(label))
}
