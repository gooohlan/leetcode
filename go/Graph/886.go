package Graph

// 将讨厌关系转换为领接表,判断是否为二分图即可
func possibleBipartition(n int, dislikes [][]int) bool {
	graph := make([][]int, n+1)
	for _, dislike := range dislikes {
		graph[dislike[0]] = append(graph[dislike[0]], dislike[1])
		graph[dislike[1]] = append(graph[dislike[1]], dislike[0])
	}

	var UNCOLORED, RED, GREEN = 0, 1, 2
	color := make([]int, n+1)

	for i := 0; i <= n; i++ {
		if color[i] == UNCOLORED { // 未访问过
			q := []int{i} // 将节点加入队列
			color[i] = RED

			for len(q) > 0 {
				node := q[0] // 取出队列头
				q = q[1:]
				nC := RED // 下一个颜色
				if color[node] == RED {
					nC = GREEN
				}

				for _, v := range graph[node] {
					if color[v] == UNCOLORED { // 未访问过
						q = append(q, v) // 加入队列
						color[v] = nC    // 染色
					} else if color[v] == color[node] {
						return false
					}
				}
			}
		}
	}

	return true
}
