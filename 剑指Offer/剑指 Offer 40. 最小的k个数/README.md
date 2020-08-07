#### [剑指 Offer 40. 最小的k个数](https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/)

输入整数数组 `arr` ，找出其中最小的 `k` 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

**示例 1：**

```
输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]
```

**示例 2：**

```
输入：arr = [0,1,2,1], k = 1
输出：[0]
```

**限制：**

- `0 <= k <= arr.length <= 10000`
- `0 <= arr[i] <= 10000`

---

快排方法：

```go
func getLeastNumbers(arr []int, k int) []int {
	n := len(arr)
	if n <= 0 || k <= 0 || k > n {
		return nil
	}
	start, end := 0, n-1
	index := partition(arr, start, end)
	for index != k-1 {
		if index > k-1 {
			end = index - 1
			index = partition(arr, start, end)
		} else {
			//注意，我就是这里犯糊涂了，浪费了一个上午的时间
			//当基准数小于k值时，我们需要将基准数搜索范围确定到右边
			start = index + 1
			index = partition(arr, start, end)
		}
	}
	// fmt.Println(arr)
	return arr[:k]
}

//快排,返回的index,是小于base的值，都在index的左边，大于base的数，都在index的右边
func partition(nums []int, start, end int) int {
	pivot := nums[start]
	for start < end {
		for start < end && nums[end] >= pivot {
			end--
		}
		nums[start] = nums[end]
		for start < end && nums[start] <= pivot {
			start++
		}
		nums[end] = nums[start]
	}
	nums[start] = pivot

	return start
}

/**
//快排,返回的index,是小于base的值，都在index的左边，大于base的数，都在index的右边
//这个方法没有将基准值移动
func partition(nums []int, left, right int) int {
	base := nums[left]
	for left < right {
		for left < right && nums[left] <= base {
			left++
		}
		for right > left && nums[right] > base {
			right--
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	//fmt.Println(nums, left)
	return left
}
**/
```

堆排：

```go
package main

import "fmt"

//快排和堆排
func getLeastNumbers(arr []int, k int) []int {
	n := len(arr)
	if n <= 0 || k <= 0 || k > n {
		return nil
	}

	nums := arr[:k]
	for i := len(nums)/2 - 1; i >= 0; i-- {
		sink(nums, i)
	}
	fmt.Println(nums)

	for i := k; i < n; i++ {
		if nums[0] <= arr[i] {
			//堆顶是最大值，最大值比arr[i]小，则说明当前arr[i]不在答案的序列中
			continue
		} else {
			nums[0] = arr[i]
			sink(nums, 0)
		}
	}

	return nums
}

//sink 下沉
func sink(nums []int, i int) {
	n := len(nums)
	for k := 2*i + 1; k < n; k = 2*k + 1 {
		if k+1 < n && nums[k] < nums[k+1] {
			k++
		}
		if nums[k] < nums[i] {
			break
		}
		//交换
		nums[i], nums[k] = nums[k], nums[i]
		i = k
	}
}

func main() {
	arr := []int{3, 1, 1, 9, 6}
	k := 2
	fmt.Println(getLeastNumbers(arr, k))
}
```

