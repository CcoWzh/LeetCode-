#### [645. 错误的集合](https://leetcode-cn.com/problems/set-mismatch/)

集合 `S` 包含从1到 `n` 的整数。不幸的是，因为数据错误，导致集合里面某一个元素复制了成了集合里面的另外一个元素的值，导致集合丢失了一个整数并且有一个元素重复。

给定一个数组 `nums` 代表了集合 `S` 发生错误后的结果。你的任务是首先寻找到重复出现的整数，再找到丢失的整数，将它们以数组的形式返回。

**示例 1:**

```
输入: nums = [1,2,2,4]
输出: [2,3]
```

**注意:**

1. 给定数组的长度范围是 [2, 10000]。
2. 给定的数组是无序的。

----

这题虽然是简单题，但是确实可以寻找更加高效的解法：

```go
func findErrorNums(nums []int) []int {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	a, b := 0, 0
	for i := 1; i <= len(nums); i++ {
		if m[i] == 2 {
			a = i
		} else if m[i] == 0 {
			b = i
		}
	}

	return []int{a, b}
}
```

对于这种数组问题，**关键点在于元素和索引是成对儿出现的，常用的方法是排序、异或、映射**。

映射的思路就是我们刚才的分析，将每个索引和元素映射起来，通过正负号记录某个元素是否被映射。

排序的方法也很好理解，对于这个问题，可以想象如果元素都被从小到大排序，如果发现索引对应的元素如果不相符，就可以找到重复和缺失的元素。

异或运算也是常用的，因为异或性质`a ^ a = 0, a ^ 0 = a`，如果将索引和元素同时异或，就可以消除成对儿的索引和元素，留下的就是重复或者缺失的元素。可以看看前文「[寻找缺失元素](http://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247484477&idx=1&sn=13834cfd618377385226c3dc598b2c28&chksm=9bd7fa35aca0732374dc34c78c276b982605892caf69cb31ad9a6c3685de5dbccac81989b195&scene=21#wechat_redirect)」，介绍过这种方法。

----

这题，还能有更高效的解法，将空间复杂度降低：

利用数组和下标的特点，把数组当成`map`：

```go
func findErrorNums(nums []int) []int {
	a, b := 0, 0 //a:重复的数字，b:缺失的数字
	n := len(nums)
	for i := 0; i < n; i++ {
		index := abs(nums[i]) - 1 //1到n的整数,不会越界
		if nums[index] < 0 {
			a = index + 1
		} else {
			//注意，这里是由于题目限制，才敢使用相反数标记以及遍历过了
			//不然，这个也得加绝对值后变负数
			nums[index] *= -1
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			b = i + 1
		}
	}
	return []int{a, b}
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}
```

参考

- [高效寻找缺失和重复的数字](https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247485050&idx=1&sn=dac757454b2df9a1291f1e8027f56c1b&chksm=9bd7f872aca07164da3e1138df630d63b41e4f071f71194069bcedeef866c1b91650a77be3d2&scene=21#wechat_redirect)



