#### [438. 找到字符串中所有字母异位词](https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/)

给定一个字符串 **s** 和一个非空字符串 **p**，找到 **s** 中所有是 **p** 的字母异位词的子串，返回这些子串的起始索引。

字符串只包含小写英文字母，并且字符串 **s** 和 **p** 的长度都不超过 20100。

**说明：**

- 字母异位词指字母相同，但排列不同的字符串。
- 不考虑答案输出的顺序。

**示例 1:**

```
输入:
s: "cbaebabacd" p: "abc"

输出:
[0, 6]

解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
```

 **示例 2:**

```
输入:
s: "abab" p: "ab"

输出:
[0, 1, 2]

解释:
起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。
```

----

滑动窗口框架：

```go
/* 滑动窗口算法框架 */
void slidingWindow(string s, string t) {
    unordered_map<char, int> need, window;
    for (char c : t) need[c]++;

    int left = 0, right = 0;
    int valid = 0; 
    while (right < s.size()) {
        // c 是将移入窗口的字符
        char c = s[right];
        // 右移窗口
        right++;
        // 进行窗口内数据的一系列更新
        ...

        /*** debug 输出的位置 ***/
        printf("window: [%d, %d)\n", left, right);
        /********************/

        // 判断左侧窗口是否要收缩
        while (window needs shrink) {
            // d 是将移出窗口的字符
            char d = s[left];
            // 左移窗口
            left++;
            // 进行窗口内数据的一系列更新
            ...
        }
    }
}
```

根据框架，写出代码：

**其中`valid`变量表示窗口中满足`need`条件的字符个数**，如果`valid`和`need.size`的大小相同，则说明窗口已满足条件，已经完全覆盖了串`T`。

**现在开始套模板，只需要思考以下四个问题**：

**1、**当移动`right`扩大窗口，即加入字符时，应该更新哪些数据？

**2、**什么条件下，窗口应该暂停扩大，开始移动`left`缩小窗口？

**3、**当移动`left`缩小窗口，即移出字符时，应该更新哪些数据？

**4、**我们要的结果应该在扩大窗口时还是缩小窗口时进行更新？

如果一个字符进入窗口，应该增加`window`计数器；如果一个字符将移出窗口的时候，应该减少`window`计数器；当`valid`满足`need`时应该收缩窗口；应该在收缩窗口的时候更新最终结果。

```go
func findAnagrams(s string, p string) []int {
	left, right := 0, 0
	window := make(map[uint8]int)
	need := make(map[uint8]int)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	valid := 0 //当valid有效时，更新答案
	res := make([]int, 0)
	for right < len(s) {
		cur := s[right]
		right++ //右移窗口
		//进行窗口内数据的一系列更新
		if _, ok := need[cur]; ok {
			window[cur]++
			if window[cur] == need[cur] {
				valid++
			}
		}
		//滑动窗口，窗口大小最大是len(p) //似乎，可以改成if
		for right-left == len(p) { //当子串是连续的时候
			if valid == len(need) {
				res = append(res, left)
			}
			d := s[left]
			//缩小窗口
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	return res
}
```







参考

- [我写了套框架，把滑动窗口算法变成了默写题](https://mp.weixin.qq.com/s/ioKXTMZufDECBUwRRp3zaA)