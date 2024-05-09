package Graph

import (
	"container/heap"
	"math"
	"sort"
)

var dir4 = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 方向数组，上下左右的坐标偏移量

// 优先队列
type point struct {
	x, y, maxDiff int
}
type hp1631 []point

func (h hp1631) Len() int              { return len(h) }
func (h hp1631) Less(i, j int) bool    { return h[i].maxDiff < h[j].maxDiff }
func (h hp1631) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp1631) Push(x interface{})   { *h = append(*h, x.(point)) }
func (h *hp1631) Pop() (v interface{}) { old := *h; *h, v = old[:len(old)-1], old[len(old)-1]; return }

func minimumEffortPath(heights [][]int) int {
	n := len(heights)
	m := len(heights[0])
	maxDiff := make([][]int, n)
	for i := range maxDiff {
		maxDiff[i] = make([]int, m)
		for j := range maxDiff[i] {
			maxDiff[i][j] = math.MaxInt
		}
	}

	maxDiff[0][0] = 0
	pq := &hp1631{{0, 0, 0}}
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(point)
		curX, curY, curDiff := cur.x, cur.y, cur.maxDiff

		if curX == n-1 && curY == m-1 { // 到达右下角,提前结束
			return curDiff
		}

		if curDiff > maxDiff[curX][curY] {
			continue
		}

		// 将 (curX, curY) 的相邻坐标装入队列
		for _, d := range dir4 {
			x, y := curX+d[0], curY+d[1]
			if x < 0 || x >= n || y < 0 || y >= m {
				continue
			}
			// 计算从 (curX, curY) 达到 (x, y) 的消耗
			diff := max(curDiff, abs(heights[x][y], heights[curX][curY]))
			if diff < maxDiff[x][y] {
				maxDiff[x][y] = diff
				heap.Push(pq, point{x, y, diff})
			}
		}
	}
	return 0
}

type edge struct {
	v, w, diff int
}

// 将所有点距离放入并查集,然后安装权重值从小到大进行排序,并依次加入并查集中,最终使得我们能够连接到右下角时退出
func minimumEffortPathUF(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	edges := make([]edge, 0)
	for i, row := range heights {
		for j, h := range row {
			id := i*m + j
			if i > 0 {
				edges = append(edges, edge{id - m, id, abs(h, heights[i-1][j])}) // [i,j]上方的体力差
			}
			if j > 0 {
				edges = append(edges, edge{id - 1, id, abs(h, heights[i][j-1])}) // [i,j]左侧的体力差
			}
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].diff < edges[j].diff
	})

	uf := NewUF(n * m)
	for _, e := range edges {
		uf.Union(e.v, e.w)          // 连接两个节点
		if uf.Connected(0, n*m-1) { // 已经连接到右下角节点
			return e.diff // 因为是取最大,edges又是按照从小到大排序的,所以放回当前节点diff即可
		}
	}

	return 0
}
