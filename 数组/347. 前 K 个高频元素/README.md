#### [347. 前 K 个高频元素](https://leetcode-cn.com/problems/top-k-frequent-elements/)

给定一个非空的整数数组，返回其中出现频率前 ***k\*** 高的元素。

 

**示例 1:**

```
输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
```

**示例 2:**

```
输入: nums = [1], k = 1
输出: [1]
```

**提示：**

- 你可以假设给定的 *k* 总是合理的，且 1 ≤ k ≤ 数组中不相同的元素的个数。
- 你的算法的时间复杂度**必须**优于 O(*n* log *n*) , *n* 是数组的大小。
- 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的。
- 你可以按任意顺序返回答案。

----

开胃菜：

手工构建小根堆

```go
func topKFrequent(nums []int, k int) []int {
	n := len(nums)
	numMap := make(map[int]int)
	for i := 0; i < n; i++ {
		numMap[nums[i]]++
	}

	size := 0
	arr := make([]int, 0)
	for i := 0; i < n; i++ {
		if size < k {
			arr = append(arr, nums[i])
			n := len(arr)
			swin(arr, n-1) //上浮
			size++
		} else { //当超过了k值后，开始更新小根堆
			//需要维护前k大的数
			if nums[i] > arr[0] { //如果新的元素的频率比堆顶端的元素大，更新
				//删除堆顶元素，将新入的元素加入堆顶，下沉
				arr[0] = nums[i]
				sink(arr, 0)
			}
		}
	}
	fmt.Println(arr)

	return arr
}

//sink 下沉
func sink(nums []int, i int) {
	n := len(nums)
	for k := 2*i + 1; k < n; k = 2*k + 1 {
		if k+1 < n && nums[k] > nums[k+1] { //比较两个左右孩子
			k++
		}
		if nums[k] > nums[i] {
			break
		}
		//交换
		nums[i], nums[k] = nums[k], nums[i]
		i = k
	}
}

//上浮
func swin(nums []int, i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if nums[parent] > nums[i] {
			nums[parent], nums[i] = nums[i], nums[parent]
		} else {
			break
		}
		i = parent
	}
}
```

将其扩展成二维数组，具体操作为：

1. 借助 哈希表 来建立数字和其出现次数的映射，遍历一遍数组统计元素的频率
2. 维护一个元素数目为 k 的最小堆
3. 每次都将新的元素与堆顶元素（堆中频率最小的元素）进行比较，（堆，用二维数组存储）
4. 如果新的元素的频率比堆顶端的元素大，则弹出堆顶端的元素，将新的元素添加进堆中
5. 最终，堆中的 k 个元素即为前 k 个高频元素

