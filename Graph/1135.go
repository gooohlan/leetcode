package Graph

import "sort"

// 先根据连接成本排序,然后再通过UF联通所有边,最终判断联通的边舒服符合要求
func minimumCost(n int, connections [][]int) int {
	var (
		total, edge int // 记录总成本和道路数量
	)
	// 排序,优先选取成本低的道路
	sort.Slice(connections, func(i, j int) bool {
		return connections[i][2] < connections[j][2]
	})

	uf := NewUF(n + 1)
	for _, connection := range connections {
		if uf.Connected(connection[0], connection[1]) { // 此条边已存在,直接跳过
			continue
		}
		// 此边不存在,连接边
		total += connection[2]
		edge++
		uf.Union(connection[0], connection[1])

		if edge == n-1 { // 已连接所有边,直接返回费用
			return total
		}
	}

	return -1 // 无法全部连通
}

type UF struct {
	n      int
	parent []int
}

func NewUF(n int) *UF {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i // 父节点指针初始指向自己
	}
	return &UF{parent: parent, n: n}
}

func (u *UF) Find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.Find(u.parent[x])
	}
	return u.parent[x]
}

func (u *UF) Union(p int, q int) {
	rootP, rootQ := u.Find(p), u.Find(q)
	if rootP == rootQ {
		return
	}

	u.parent[rootQ] = rootP
}

func (u *UF) Connected(p int, q int) bool {
	return u.Find(p) == u.Find(q)
}
