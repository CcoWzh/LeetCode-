#### [39. 组合总和](https://leetcode-cn.com/problems/combination-sum/)

给定一个**无重复元素**的数组 `candidates` 和一个目标数 `target` ，找出 `candidates` 中所有可以使数字和为 `target` 的组合。

`candidates` 中的数字可以无限制重复被选取。

**说明：**

- 所有数字（包括 `target`）都是正整数。
- 解集不能包含重复的组合。 

**示例 1:**

```
输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
```

**示例 2:**

```
输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
```

----

回溯法

- [Leecode题解 CombinationSum](https://mp.weixin.qq.com/s?__biz=MzU3NDgxMzI0Mw==&mid=2247485660&idx=2&sn=e8a7e525b477e96ec737f48ae455cbc8&chksm=fd2ded88ca5a649e21f74b2efda8f88f2d3d88af5ac248a443f031de2858da12a851e04c1ac5&scene=126&sessionid=1588556524&key=b2f7a3ef4ee6a5af8153ca9f3c145508164efaeb0c9f7f96b0abdc4ea513f9268ba8989e6e1c4a190545c3ff8b038ce1728cfce7bd8383f71db2597f1f04a5cc4fc522cb56ceff60ef567eff3b24200d&ascene=1&uin=OTg2NTk3NTA2&devicetype=Windows+10+x64&version=62090070&lang=zh_CN&exportkey=A3I4GQBCata5qubciG6u0%2Bo%3D&pass_ticket=Rzy3RbyicL%2Bnp0qUXlTY3yKFrC9v65uiZAMvB5ty3BzLAaDH%2FWQX%2FoedD5LGQAG0)

