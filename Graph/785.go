package Graph

func isBipartite(graph [][]int) bool {
	var (
		visited = make(map[int]bool)
		color   = make([]bool, len(graph))
		ok      = true
		dfs     func(int)
	)

	dfs = func(node int) {
		visited[node] = true
		for _, v := range graph[node] {
			if !visited[v] { // 未访问过
				// 给当前节点涂上相反的颜色
				color[v] = !color[node]
				// 继续给V节点着色
				dfs(v)
			} else {
				// 访问过的节点,如果颜色相同表示非二色图
				if color[v] == color[node] {
					ok = false
					return
				}
			}
		}
	}

	for i := 0; i < len(graph) && ok; i++ {
		if !visited[i] {
			dfs(i)
		}
	}
	return ok
}

func isBipartiteBFS(graph [][]int) bool {
	var (
		visited = make(map[int]bool)
		color   = make([]bool, len(graph))
	)
	for i := 0; i < len(graph); i++ {
		if !visited[i] {
			var queue []int
			queue = append(queue, i)
			for len(queue) > 0 {
				node := queue[0]
				queue = queue[1:]

				for _, v := range graph[node] {
					if !visited[v] { // 未访问过
						color[v] = !color[node]  // 给当前节点涂上相反的颜色
						visited[v] = true        // 标记V访问过
						queue = append(queue, v) // v加入队列
					} else {
						// 访问过的节点,如果颜色相同表示非二色图
						if color[v] == color[node] {
							return false
						}
					}
				}
			}
		}
	}
	return true
}
