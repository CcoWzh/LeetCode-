#### [剑指 Offer 45. 把数组排成最小的数](https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/)

输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。

**示例 1:**

```
输入: [10,2]
输出: "102"
```

**示例 2:**

```
输入: [3,30,34,5,9]
输出: "3033459"
```

**提示:**

- `0 < nums.length <= 100`

**说明:**

- 输出结果可能非常大，所以你需要返回一个字符串而不是整数
- 拼接起来的数字可能会有前导 0，最后结果不需要去掉前导 0

---

回溯法（ 超出时间限制）：

这个是可行的，编程正确，但是会超出时间限制。

```go
package main

import (
	"fmt"
	"strconv"
)

//时间超出限制
func minNumber(nums []int) string {
	n := len(nums)
	var result string
	if n == 0 {
		return result
	}

	backTrack(nums, "", 0, n, &result)
	return result
}

//回溯，nums：可选择的列表，path已经做过的选择，count选择的次数，length 总长度
func backTrack(nums []int, path string, count, length int, result *string) {
	if count == length {
		*result = compareStr(*result, path)
	}

	for i := 0; i < len(nums); i++ {
		temp := path
		path += strconv.Itoa(nums[i])

		backTrack(deleteNums(nums, i), path, count+1, length, result)

		path = temp
	}
}

func deleteNums(nums []int, index int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	//不能直接负责，因为切片是浅拷贝，拷贝的是数据地址
	res := make([]int, 0)
	res = append(res, nums[:index]...)
	res = append(res, nums[index+1:]...)
	return res
}

//假设，a,b 的长度相等,返回最小值
func compareStr(a, b string) string {
	n, m := len(a), len(b)
	if n == 0 {
		return b
	} else if m == 0 {
		return a
	}
	for i := 0; i < n; i++ {
		if a[i] == b[i] {
			continue
		}
		if a[i] > b[i] {
			return b
		} else if a[i] < b[i] {
			return a
		}
	}
	return a
}

func main() {
	nums := []int{3, 30, 0, 5, 0}
	fmt.Println(minNumber(nums))
}

```

排序（比较大小）