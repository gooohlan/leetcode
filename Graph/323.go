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

	dfs = func(cur int) {
		visited[cur] = true
		for _, v := range graph[cur] {
			if !visited[v] {
				dfs(v)
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

func countComponentsBfs(n int, edges [][]int) int {
	var (
		graph   = make([][]int, n)
		visited = make(map[int]bool) // 记录遍历过的节点，防止走回头路
	)

	for _, item := range edges {
		graph[item[1]] = append(graph[item[1]], item[0])
		graph[item[0]] = append(graph[item[0]], item[1])
	}

	var res int
	for i := 0; i < n; i++ {
		if !visited[i] {
			res++
			queue := []int{i}
			for len(queue) > 0 {
				cur := queue[0]
				queue = queue[1:]

				for _, v := range graph[cur] {
					if !visited[v] {
						visited[v] = true
						queue = append(queue, v)
					}
				}
			}
		}
	}

	return res
}

func countComponentsBfsUF(n int, edges [][]int) int {
	var (
		count  = n
		parent = make([]int, n) // 父节点
		find   func(int) int
		union  func(int, int)
		// connected func(int, int) bool
	)

	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union = func(p int, q int) {
		rootP, rootQ := find(p), find(q)
		if rootP == rootQ {
			return
		}

		parent[rootQ] = rootP
		count--
	}

	for i := 0; i < n; i++ {
		parent[i] = i // 父节点指针初始指向自己
	}

	for _, edge := range edges {
		union(edge[0], edge[1])
	}

	return count
}
