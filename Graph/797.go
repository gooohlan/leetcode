package Graph

func allPathsSourceTarget(graph [][]int) [][]int {
	var path []int
	var res [][]int
	var travers func(s int)
	travers = func(s int) {
		path = append(path, s) // 添加节点s 到路径
		if s == len(graph)-1 { // 到达终点
			item := make([]int, len(path))
			copy(item, path)
			res = append(res, item)
		}

		// 遍历相邻的每个节点
		for _, v := range graph[s] {
			travers(v)
		}
		path = path[:len(path)-1]
	}

	travers(0)
	return res
}
