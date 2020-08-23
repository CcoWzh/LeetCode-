package main

func nextPermutation(nums []int) {
	n := len(nums)
	if n == 0 {
		return
	}
	//找到第一个降序的数字,即，从后向前查找第一个相邻升序的元素对
	index := n - 2 //从倒数第二个数出发
	for index >= 0 {
		if nums[index] < nums[index+1] {
			break
		}
		index--
	}
	//从后向前查找第一个满足 A[i] < A[k] 的 k,而不是从前往后查找
	k := n - 1
	if index >= 0 { //判断，如果没有降序的数字的话，说明不存在下一个更大的排列
		for nums[index] >= nums[k] {
			k--
		}
		//交换数字
		nums[index], nums[k] = nums[k], nums[index]
	}

	//从index的位置，开始排序。可以断定这时 [index,end) 必然是降序，逆置 [j,end)，使其升序
	for i, j := index+1, n-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	//fmt.Println(nums)
}

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums) //原地修改
}
