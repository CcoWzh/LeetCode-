#### [31. 下一个排列](https://leetcode-cn.com/problems/next-permutation/)

实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须**[原地](https://baike.baidu.com/item/原地算法)**修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
`1,2,3` → `1,3,2`
`3,2,1` → `1,2,3`
`1,1,5` → `1,5,1`

----

我们希望

- 下一个数比当前数大，这样才满足“下一个排列”的定义。
- 我们还希望下一个数增加的幅度尽可能的小。

算法过程，标准的“下一个排列”算法可以描述为：

1. 从后向前查找第一个相邻升序的元素对 (i,j)，满足 A[i] < A[j]。此时 [j,end) 必然是降序
2. 在 [j,end) 从后向前查找第一个满足 A[i] < A[k] 的 k。A[i]、A[k] 分别就是上文所说的「小数」、「大数」
3. 将 A[i] 与 A[k] 交换
4. 可以断定这时 [j,end) 必然是降序，逆置 [j,end)，使其升序
5. 如果在步骤 1 找不到符合的相邻元素对，说明当前 [begin,end) 为一个降序顺序，则直接跳到步骤 4

![image.png](https://pic.leetcode-cn.com/e56a66ed318d1761cd8c8f9d1521f82a30c71ecc84f551912b90d8fe254c8f3d-image.png)

原版：

```go
func nextPermutation(nums []int) {
	n := len(nums)
	if n == 0 {
		return
	}
	//找到第一个降序的数字
	index := n - 2
	for index >= 0 {
		if nums[index] < nums[index+1] {
			break
		}
		index--
	}

	cIndex, min := index+1, nums[index+1]
	for i := index + 1; i < n; i++ {
		if nums[i] > nums[index] && min > nums[i] {
			min = nums[i]
			cIndex = i
		}
		if nums[i] < nums[index] {
			break
		}
	}

	nums[index], nums[cIndex] = nums[cIndex], nums[index]
	//从index的位置，开始排序,可以断定这时 [cIndex,end) 必然是降序，逆置 [j,end)，使其升序
	for i := index+1; i < n; i++ {
		for j := i; j < n-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	fmt.Println(nums)
}
```

修改后：

```go
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
```

参考

- [下一个排列](https://leetcode-cn.com/problems/next-permutation/solution/xia-yi-ge-pai-lie-suan-fa-xiang-jie-si-lu-tui-dao-/)