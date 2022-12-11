package WeeklyContest

import (
	"math"
	"sort"
)

// 构建无向图
// 然后遍历节点
// 每个节点前K个大于0的节点和+当前节点值,为当前节点最大星和
func maxStarSum(vals []int, edges [][]int, k int) int {
	var (
		graph = make([][]int, len(vals))
		res   = math.MinInt
	)

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	for i := 0; i < len(vals); i++ { // 遍历节点
		temp := make([]int, 0)
		for _, v := range graph[i] {
			temp = append(temp, vals[v])
		}
		sort.Slice(temp, func(i, j int) bool { // 当前节点可访问节点值倒序排序
			return temp[i] > temp[j]
		})

		kk := len(temp)
		if kk > k {
			kk = k
		}
		sum := vals[i]
		for j := 0; j < kk; j++ { // 前K个节点不为0的和为当前节点最大星和
			if temp[j] <= 0 {
				break
			}
			sum += temp[j]
		}
		res = max(sum, res)
	}

	return res
}
