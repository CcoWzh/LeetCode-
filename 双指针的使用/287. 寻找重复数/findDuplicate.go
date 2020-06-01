package main

func findDuplicate(nums []int) int {
	fast, slow := 0, 0

	for fast < len(nums) {
		fast = nums[nums[fast]]
		slow = nums[slow]
		//fmt.Println(fast,slow,nums[slow])
		if fast == slow {
			break
		}
	}
	//fmt.Println(fast, slow)

	slow = 0

	for nums[slow] != nums[fast] {
		fast = nums[fast]
		slow = nums[slow]
	}

	//fmt.Println(nums[fast])

	return nums[fast]

}

func main() {
	nums := []int{3, 1, 3, 4, 2}
	findDuplicate(nums)
}
