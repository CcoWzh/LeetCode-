package main

import "fmt"

/**
990. 等式方程的可满足性
 */
func equationsPossible(equations []string) bool {
	n := len(equations)
	if n == 0 {
		return false
	}
	uf := UF{}
	uf.newUF(26)
	//连接==的节点
	for i := 0; i < n; i++ {
		if equations[i][1] == '=' {
			x := int(equations[i][0] - 'a')
			y := int(equations[i][3] - 'a')
			uf.union(x, y)
		}
	}
	//！= 的节点不能影响== 的情况
	for i := 0; i < n; i++ {
		if equations[i][1] == '!' {
			x := int(equations[i][0] - 'a')
			y := int(equations[i][3] - 'a')
			if uf.connected(x, y) {
				return false
			}
		}
	}
	return true
}

func main() {
	equations := []string{"a==b", "b!=c", "c==a"}
	fmt.Println(equationsPossible(equations))
}

type UF struct {
	count  int   // 记录连通分量个数
	parent []int // 存储若干棵树
	size   []int // 记录树的“重量”
}

//初始化
func (uf *UF) newUF(n int) {
	uf.count = n
	uf.parent = make([]int, n)
	uf.size = make([]int, n)
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = i
	}
}

//连通两个节点（树）
func (uf *UF) union(p, q int) {
	i, j := uf.find(p), uf.find(q)
	if i == j {
		return
	}
	if uf.size[i] < uf.size[j] {
		uf.parent[i] = j
		uf.size[j] += uf.size[i]
	} else {
		uf.parent[j] = i
		uf.size[i] += uf.size[j]
	}
	uf.count--
}

//寻找根节点
func (uf *UF) find(p int) int {
	for p != uf.parent[p] {
		p = uf.parent[p]
	}
	return p
}

//判断是否连通
func (uf *UF) connected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}
