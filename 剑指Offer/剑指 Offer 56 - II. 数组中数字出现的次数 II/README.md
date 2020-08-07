#### [剑指 Offer 56 - II. 数组中数字出现的次数 II](https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/)

在一个数组 `nums` 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。

**示例 1：**

```
输入：nums = [3,4,3,3]
输出：4
```

**示例 2：**

```
输入：nums = [9,1,7,9,7,9,7]
输出：1
```

**限制：**

- `1 <= nums.length <= 10000`
- `1 <= nums[i] < 2^31`

---

思路：

```
//将每个数字的每一位都相加，这样，出现3次的数字，其位数之和一定是3的倍数，即
//如果把所有出现3次的数字的二进制表示的每一位都分别加起来，那么每一位的和都能被3整除
//不能被3整除的部分，就说明，是只出现一次
```

这里，我们用到了一个技巧：

二进制转十进制的编程方法：

```go
func twoToten(bit []int) int {
	result := 0
	for i := 0; i < len(bit); i++ {
		result = result << 1
		result += bit[i]
	}
	return result
}

func main() {
	fmt.Println(twoToten([]int{1,1,0}))
}
输出：
6
```

这里，其实就是普通的转换方法：

```mariadb
1101.01（2）=1*2^3+1*2^2+0*2^1+1*2^0+0*2-1+1*2-2=1+0+4+8+0+0.25=13.25（10）
```



![img](https://iknow-pic.cdn.bcebos.com/622762d0f703918f182e61b05c3d269758eec4c9?x-bce-process=image/resize,m_lfit,w_600,h_800,limit_1)

只不过，这个技巧在于，每一次`for`循环，都在原来相加的基础上，乘以2

