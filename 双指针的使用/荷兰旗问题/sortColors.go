package main

import "fmt"

/**
荷兰旗问题又称三色排序，或者彩虹排序
比如给你一个只含有三个数字的数组：312312312231111122113，
要求按照 1 2 3 的顺序排好，即：111111111222222222223333333333
（请不要真的去数数，认真你就输了）

具体说来，用 i, j, k 这三个指针分一下：
[0, i): 存 1
[i, j): 存 2
[j, k]: 未排序区间
(k, array.length-1]: 存 3
这里 j 放在未排序区间的左边和右边都行，但基本上大家都是放左边，
 */
func sortColors(nums []int) []int {
	i, j := 0, 0
	k := len(nums) - 1

	for j <= k {
		switch nums[j] {
		case 1:
			nums[i], nums[j] = nums[j], nums[i]
			j++
			i++
		case 2:
			j++
		case 3:
			nums[k], nums[j] = nums[j], nums[k]
			//注意这里就不能 j++ 了，因为新换回来的元素还没瞧呢，不知道它是几。
			k--
		default:
			fmt.Println("无效输入......")
		}
	}

	return nums
}

func main() {
	nums := []int{1, 2, 3, 2, 3, 1, 2}
	fmt.Println(sortColors(nums))
}
