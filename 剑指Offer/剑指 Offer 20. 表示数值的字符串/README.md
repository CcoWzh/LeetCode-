#### [剑指 Offer 20. 表示数值的字符串](https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/)

请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。例如，字符串"+100"、"5e2"、"-123"、"3.1416"、"0123"都表示数值，但"12e"、"1a3.14"、"1.2.3"、"+-5"、"-1E-16"及"12e+5.4"都不是。

注意：本题与主站 65 题相同：https://leetcode-cn.com/problems/valid-number/

---

使用标志位，来区分该符号之前是否出现过：

- `pointFlag := 0`
- `eFlag := 0`
- `numFlag := 0`

```go
func isNumber(s string) bool {
	//注意，需要预处理，去除s首位空格
	s = strings.TrimSpace(s)
	n := len(s)
	if n == 0 {
		return false
	}
	//3个标志位，判断对应的符号是否出现过
	pointFlag := 0
	eFlag := 0
	numFlag := 0
	for i := 0; i < n; i++ {
		cur := int(s[i] - '0')
		if cur >= 0 && cur <= 9 {
			numFlag = 1
		} else if s[i] == '.' && pointFlag == 0 && eFlag == 0 {
			//小数点只能出现一次，且只能出现在e前面
			pointFlag = 1
		} else if (s[i] == 'e' || s[i] == 'E') && eFlag == 0 && numFlag == 1 {
			//e前面必须得有数字，且只能出现一次
			eFlag = 1
			numFlag = 0
		} else if s[i] == '+' || s[i] == '-' {
			//正负号必须出现在开头或者e后面，且后面必须跟有数字，1e+
			if (numFlag == 0 && i == 0) || (eFlag == 1 && numFlag == 0 && i+1 != n) {
				continue
			} else {
				return false
			}
		} else {
			return false
		}
	}
	//返回的是numFlag！！！
	return numFlag == 1
}
```

