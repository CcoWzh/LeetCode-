#### [283. 移动零](https://leetcode-cn.com/problems/move-zeroes/)

给定一个数组 `nums`，编写一个函数将所有 `0` 移动到数组的末尾，同时保持非零元素的相对顺序。

**示例:**

```
输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
```

**说明**:

1. 必须在原数组上操作，不能拷贝额外的数组。
2. 尽量减少操作次数。

----

如果不考虑条件限制，还是很简单的：

```go
func moveZeroes(nums []int) {
	n := len(nums)
	cc := make([]int, 0)
	//遇到是0的，直接跳过
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			continue
		}
		cc = append(cc, nums[i])
	}
	//在cc数组后补0
	for i := len(cc); i < n; i++ {
		cc = append(cc, 0)
	}
	nums = cc //复制
	fmt.Println(nums)
}
```

但是，题目要求不能使用额外的数组，需要原地操作。

这里，我们可以使用**双指针**

```
func moveZeroes(nums []int) {
	n := len(nums)
	fast, slow := 0, 0
	for fast < n {
		if nums[fast] == 0 {
			fast++
		} else {
			nums[slow] = nums[fast]
			fast++
			slow++
		}
	}

	for i := slow; i < n; i++ {
		nums[i] = 0
	}
	fmt.Println(nums)
}
```

