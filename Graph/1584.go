package Graph

import (
	"sort"
)

// 此题与1135基本一致,不同的是边的权重需要提前计算
// 这里为了方便Union-Find算法,我们把prints的索引代表坐标点,两个坐标点的距离作为权重题目就与1135完全一样了
func minCostConnectPoints(points [][]int) int {
	n := len(points)
	var edges [][]int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			xi, yi := points[i][0], points[i][1] // 第i个点
			xj, yj := points[j][0], points[j][1] // 第j个点

			// 将i,j两个点的索引以及权重加入新数组
			edges = append(edges, []int{
				i, j, abs(xi, xj) + abs(yi, yj),
			})
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	// 后序和1135一样
	uf := NewUF(n)
	var total int
	for _, edge := range edges {
		if uf.Connected(edge[0], edge[1]) {
			continue
		}

		total += edge[2]
		uf.Union(edge[0], edge[1])
		n--
		if n == 1 { // 已经选择了N条边,后序选择的都会是已经遍历过的,提前跳出循环
			break
		}
	}
	return total
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
