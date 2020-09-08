package main

func sortColors(nums []int) {
	i, j := 0, 0
	k := len(nums) - 1
	for j <= k {
		if nums[j] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else if nums[j] == 1 {
			j++
		} else {
			nums[j], nums[k] = nums[k], nums[j]
			k--
			//j++ //注意这里就不能 j++ 了，因为新换回来的元素还没瞧呢，不知道它是几。
		}
	}
	//fmt.Println(nums)
}

func main() {
	nums := []int{}
	sortColors(nums)
}
