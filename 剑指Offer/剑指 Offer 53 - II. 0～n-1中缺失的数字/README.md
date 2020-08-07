#### [剑指 Offer 53 - II. 0～n-1中缺失的数字](https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof/)

一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。 

**示例 1:**

```
输入: [0,1,3]
输出: 2
```

**示例 2:**

```
输入: [0,1,2,3,4,5,6,7,9]
输出: 8
```

**限制：**

```
1 <= 数组长度 <= 10000
```

---

方法一，使用等差数列：

```go
func missingNumber(nums []int) int {
	n := len(nums)
	sum1 := (n + 0) * (n + 1) / 2
	sum2 := 0
	for _, v := range nums {
		sum2 += v
	}

	return sum1 - sum2
}
```

方法二，二分法：

题目转换成：寻找第一个，数组下标和数字不一致的数。

也就是用二分法，寻找左区间或者右区间的问题

```go
func missingNumber(nums []int) int {
	left, right := 0, len(nums)
	if right <= 0 {
		return 0
	}

	for left < right {
		mid := (left + right) / 2
		//当前数字等于数组下标
		if nums[mid] == mid {
			left = mid + 1
		} else if nums[mid] != mid {
			right = mid
		}
	}

	return left
}
```

参考：[我作了首诗，保你闭着眼睛也能写对二分查找](https://mp.weixin.qq.com/s/M1KfTfNlu4OCK8i9PSAmug)

方法三，位运算：

如果题目是没有排序好的数组呢？这时候怎么办？可以用等差数列求和，也可以使用位运算。

如何找这个落单的数字呢，**只要把所有的元素和索引做异或运算，成对儿的数字都会消为 0，只有这个落单的元素会剩下**，也就达到了我们的目的。

```go
func missingNumber(nums []int) int {
	n := len(nums)
	result := 0
	if n <= 0 {
		return result
	}
	result ^= n
	for i := 0; i < n; i++ {
		result ^= i ^ nums[i]
	}

	return result
}
```


