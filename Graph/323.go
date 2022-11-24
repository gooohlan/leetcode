package Graph

// dfs 建立领接表然后记录访问向量
func countComponents(n int, edges [][]int) int {
	var (
		graph   = make([][]int, n)
		visited = make(map[int]bool) // 记录遍历过的节点，防止走回头路
		dfs     func(int)
	)

	for _, item := range edges {
		graph[item[1]] = append(graph[item[1]], item[0])
		graph[item[0]] = append(graph[item[0]], item[1])
	}

	dfs = func(v int) {
		visited[v] = true
		for _, i := range graph[v] {
			if !visited[i] {
				dfs(i)
			}
		}
	}

	var res int
	for i := 0; i < n; i++ {
		if !visited[i] {
			res++
			dfs(i)
		}
	}

	return res
}
