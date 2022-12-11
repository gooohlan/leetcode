package WeeklyContest

import (
	"container/heap"
	"sort"
)

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// 使用并集查,分别排序节点和queries, 从小到大依次查找
func maxPoints(grid [][]int, queries []int) []int {
	m := len(grid)
	n := len(grid[0])
	mn := m * n

	// 并查集模版
	fa := make([]int, mn)
	size := make([]int, mn)
	for i := range fa {
		fa[i] = i
		size[i] = 1
	}
	var (
		find  func(int) int
		union func(int, int)
	)
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	union = func(from int, to int) {
		from = find(from)
		to = find(to)
		if from != to {
			fa[from] = to
			size[to] += size[from]
		}
	}

	// 矩阵元素从小到大排序，方便离线
	type tuple struct {
		val, i, j int
	}
	a := make([]tuple, 0, mn)
	for i, row := range grid {
		for j, val := range row {
			a = append(a, tuple{val, i, j})
		}
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].val < a[j].val
	})

	// 查询的下标按照查询值从小到大排序，方便离线
	id := make([]int, len(queries))
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool { return queries[id[i]] < queries[id[j]] })

	res := make([]int, len(queries))
	j := 0
	for _, i := range id {
		q := queries[i]
		for ; j < mn && a[j].val < q; j++ {
			x, y := a[j].i, a[j].j
			// 枚举上下左右,值小于q可以连接
			for _, dir := range dirs {
				x2, y2 := x+dir.x, y+dir.y
				if 0 <= x2 && x2 < m && 0 <= y2 && y2 < n && grid[x2][y2] < q {
					union(x*n+y, x2*n+y2) // 连接符合条件的点
				}
			}
		}

		if grid[0][0] < q {
			res[i] = size[find(0)]
		}
	}
	return res
}

// 使用最小堆,从左上角向外搜索,对每个接单镜像询问,如果堆顶元素小于询问值,这弹出堆定,继续搜索
func maxPoints2(grid [][]int, queries []int) []int {
	m := len(grid)
	n := len(grid[0])

	// 查询的下标按照查询值从小到大排序，方便查找
	id := make([]int, len(queries))
	for i := range id {
		id[i] = i
	}
	sort.Slice(id, func(i, j int) bool { return queries[id[i]] < queries[id[j]] })

	res := make([]int, len(queries))
	// 初始化小顶堆
	h := hp{{grid[0][0], 0, 0}}
	grid[0][0] = 0 // grid原地更改,标识已访问过,充当vis
	cnt := 0
	for _, i := range id {
		q := queries[i]
		// 堆不为空且堆顶小于当前查找元素,可继续查找
		for len(h) > 0 && h[0].val < q {
			cnt++
			p := heap.Pop(&h).(tuple)
			for _, dir := range dirs {
				x, y := p.i+dir.x, p.j+dir.y
				if 0 <= x && x < m && 0 <= y && y < n && grid[x][y] > 0 { // 大于0表示没访问过
					heap.Push(&h, tuple{grid[x][y], x, y})
					grid[x][y] = 0 // 标记已访问
				}
			}
		}
		res[i] = cnt
	}

	return res
}

// 最小堆
type tuple struct{ val, i, j int }
type hp []tuple

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].val < h[j].val }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(tuple)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
