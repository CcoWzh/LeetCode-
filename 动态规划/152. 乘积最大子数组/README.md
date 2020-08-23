#### [152. 乘积最大子数组](https://leetcode-cn.com/problems/maximum-product-subarray/)

给你一个整数数组 `nums` ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

**示例 1:**

```
输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
```

**示例 2:**

```
输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
```

---

动态规划

由于存在负数，那么会导致最大的变最小的，最小的变最大的。因此还需要维护当前最小值`min`，

```go
dpMax[i] = max(nums[i], max(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i]))
dpMin[i] = min(nums[i], min(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i]))
```

当负数出现时则`max`与`min`进行交换再进行下一步计算

