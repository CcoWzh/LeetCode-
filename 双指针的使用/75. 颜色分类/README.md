#### [75. 颜色分类](https://leetcode-cn.com/problems/sort-colors/)

给定一个包含红色、白色和蓝色，一共 *n* 个元素的数组，**[原地](https://baike.baidu.com/item/原地算法)**对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

**注意:**
不能使用代码库中的排序函数来解决这道题。

**示例:**

```
输入: [2,0,2,1,1,0]
输出: [0,0,1,1,2,2]
```

**进阶：**

- 一个直观的解决方案是使用计数排序的两趟扫描算法。
  首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
- 你能想出一个仅使用常数空间的一趟扫描算法吗？

---

使用三指针，直接一趟扫描：

```go
func sortColors(nums []int) {
	i, j := 0, 0
	k := len(nums) - 1
	for j <= k {
		if nums[j] == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else if nums[j] == 1 {
			j++
		} else {
			nums[j], nums[k] = nums[k], nums[j]
			k--
			//j++ //注意这里就不能 j++ 了，因为新换回来的元素还没瞧呢，不知道它是几。
		}
	}
	//fmt.Println(nums)
}
```

附录：

-----

[算法题] 荷兰国旗问题

荷兰旗问题又称三色排序，或者彩虹排序，

![img](https://mmbiz.qpic.cn/mmbiz_png/og5r7PHGHojV6Sw9k2JVCnwNukNibOkrHxFa7ibhP12NhPNjkA9769E4EkwHiawqVKME0EheTAo2G0At9uhtWpib1A/640?wx_fmt=png&tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)

因为荷兰旗就三种颜色嘛，那这道题的问题就是给你三种颜色，按照给定的顺序排好。

当然了，题目的问法各种各样，有的给数字，有的给字母，但本质都是一样的。

比如给你一个只含有三个数字的数组：312312312231111122113，
要求按照 1 2 3 的顺序排好，即：111111111222222222223333333333

----

三指针解决问题：

```go
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
```

参考

- [面试算法：荷兰旗问题](https://mp.weixin.qq.com/s/OypQ-mzIqJQ_3pohi9wBGQ)

