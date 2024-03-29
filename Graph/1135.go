package Graph

import (
	"container/heap"
	"sort"
)

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

// Prim 算法解法
func minimumCostPrim(n int, connections [][]int) int {
	var (
		total, edge int                    // 记录总成本和道路数量
		graph       = make([][][]int, n+1) // graph[s] 记录节点 s 所有相邻的边, 三元组 int[]{from, to, weight} 表示一条边
		inMst       = make(map[int]bool)   // 记录边是否出现过
		cut         func(int)
		pq          PriorityQueue
	)

	cut = func(s int) {
		for _, v := range graph[s] {
			if inMst[v[1]] {
				continue
			}
			heap.Push(&pq, v)
		}
	}

	for _, connection := range connections {
		graph[connection[0]] = append(graph[connection[0]], []int{connection[0], connection[1], connection[2]})
		graph[connection[1]] = append(graph[connection[1]], []int{connection[1], connection[0], connection[2]})
	}

	inMst[1] = true
	cut(1)
	for pq.Len() > 0 {
		cur := heap.Pop(&pq).([]int)
		to, weight := cur[1], cur[2]
		if inMst[to] {
			// 节点 to 已经在最小生成树中，跳过
			// 否则这条边会产生环
			continue
		}
		total += weight
		edge++
		inMst[to] = true
		cut(to)
		if edge == n-1 { // 已连接所有边,直接返回费用
			return total
		}
	}
	return total
}

type PriorityQueue [][]int

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i][2] < pq[j][2]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.([]int))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
