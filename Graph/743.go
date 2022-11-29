package Graph

import (
	"container/heap"
	"math"
)

func networkDelayTime(times [][]int, n int, k int) int {
	type edge struct {
		to, time int // to 目标, time 耗时
	}
	// 构建图
	graph := make([][]edge, n)
	for _, time := range times {
		graph[time[0]-1] = append(graph[time[0]-1], edge{time[1] - 1, time[2]})
	}
	dist := make([]int, n)
	for i := range dist {
		dist[i] = math.MaxInt
	}
	dist[k-1] = 0 // 开始节点到开始节点的耗时为0
	pq := &hp{{0, k - 1}}

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(pair)
		if dist[cur.id] < cur.d { // 当前节点距离大于已知节点最大直接跳过
			continue
		}

		for _, e := range graph[cur.id] {
			if d := dist[cur.id] + e.time; d < dist[e.to] { // 小于当前已知最小距离,更新距离值
				dist[e.to] = d
				heap.Push(pq, pair{d, e.to})
			}
		}
	}

	res := 0
	for _, d := range dist {
		if d == math.MaxInt {
			return -1
		}
		res = max(res, d)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type pair struct {
	d, id int // d 距离, id 图节点id
}
type hp []pair

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i].d < h[j].d
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *hp) Pop() (v interface{}) {
	old := *h
	*h, v = old[:len(old)-1], old[len(old)-1]
	return
}
