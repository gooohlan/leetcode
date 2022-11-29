package Graph

import "container/heap"

type prob struct {
	node    int
	maxProb float64
}
type hp1514 []prob

func (h hp1514) Len() int              { return len(h) }
func (h hp1514) Less(i, j int) bool    { return h[i].maxProb > h[j].maxProb }
func (h hp1514) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp1514) Push(x interface{})   { *h = append(*h, x.(prob)) }
func (h *hp1514) Pop() (v interface{}) { old := *h; *h, v = old[:len(old)-1], old[len(old)-1]; return }

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	type edge struct {
		to   int     // 目标
		prob float64 // 概率
	}
	// 构建图
	graph := make([][]edge, n)
	for i := range edges {
		graph[edges[i][0]] = append(graph[edges[i][0]], edge{edges[i][1], succProb[i]})
		graph[edges[i][1]] = append(graph[edges[i][1]], edge{edges[i][0], succProb[i]})
	}

	dist := make([]float64, n)
	for i := range dist {
		dist[i] = -1
	}
	dist[start] = 1 // 连接自身的概率为1
	pq := &hp1514{{start, 1}}

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(prob)
		curNode := cur.node
		curProb := cur.maxProb

		if curNode == end { // 遇到终点,提前结束
			return curProb
		}

		if curProb < dist[curNode] {
			// 已有更大概率的路径
			continue
		}

		for _, e := range graph[curNode] {
			nextNode := e.to
			nextProb := dist[curNode] * e.prob
			if dist[nextNode] < nextProb {
				dist[nextNode] = nextProb
				heap.Push(pq, prob{nextNode, nextProb})
			}
		}
	}
	return 0
}
