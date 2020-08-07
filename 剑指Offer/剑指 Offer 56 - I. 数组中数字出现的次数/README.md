#### [剑指 Offer 56 - I. 数组中数字出现的次数](https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/)

一个整型数组 `nums` 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。

**示例 1：**

```
输入：nums = [4,1,4,6]
输出：[1,6] 或 [6,1]
```

**示例 2：**

```
输入：nums = [1,2,10,4,1,4,3,3]
输出：[2,10] 或 [10,2]
```

**限制：**

- `2 <= nums.length <= 10000`

----



````go
func singleNumbers(nums []int) []int {
	n := len(nums)
	ruler := 0

	for i := 0; i < n; i++ {
		ruler ^= nums[i]
	}

	index := 0
	for ruler&1 == 0 && ruler != 0 {
		ruler = ruler >> 1
		index++
	}
	one, other := make([]int, 0), make([]int, 0)

	for i := 0; i < n; i++ {
		if isBit(index, nums[i]) {
			one = append(one, nums[i])
		} else {
			other = append(other, nums[i])
		}
	}

	a, b := 0, 0
	for i := 0; i < len(one); i++ {
		a ^= one[i]
	}
	for i := 0; i < len(other); i++ {
		b ^= other[i]
	}

	return []int{a, b}
}

func isBit(index int, num int) bool {
	num = num >> uint(index)
	return num&1 == 1
}
````

