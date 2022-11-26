package Graph

// dfs
// 根据edges创建领接表, 此题要求问不存在环切只有通过一个节点可以联通所有节点
// 从节点0出发遍历领接表,出现环则直接返回false
// 此处我们引入一个细节,如果这个领接表是个数,那么他的边树应该为n-1,否则必存在环
// 最后判断是否存在未遍历节点
func validTreeDFS(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}
	var (
		graph   = make([][]int, n)
		visited = make(map[int]bool, n)
		dfs     func(int)
	)

	// 创建领接表
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	dfs = func(node int) {
		visited[node] = true

		for _, v := range graph[node] {
			if !visited[v] {
				dfs(v)
			}
		}
	}

	dfs(0)
	for i := 0; i < n; i++ {
		if !visited[i] { // 没被遍历过则表示不可达,出现多颗树
			return false
		}
	}

	return true
}

// BFS
// 解法思路与上述相同,只是修改了遍历方式
func validTreeBFS(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}
	var (
		graph   = make([][]int, n)
		visited = make(map[int]bool, n)
	)

	// 创建领接表
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	q := []int{0}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		visited[cur] = true
		for _, v := range graph[cur] {
			if !visited[v] {
				q = append(q, v)
			}
		}
	}

	for i := 0; i < n; i++ {
		if !visited[i] { // 没被遍历过则表示不可达,出现多颗树
			return false
		}
	}

	return true
}

// UF
// 使用并查集的思路,如果两个节点本身就处于同一连通分量重,再添加这条边必产生环,反之不会产生环
func validTreeUF(n int, edges [][]int) bool {
	var (
		parent = make([]int, n) // 父节点
		find   func(int) int
		union  func(int, int)
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
		n--
	}

	for i := 0; i < n; i++ {
		parent[i] = i // 父节点指针初始指向自己
	}

	for _, edge := range edges {
		if find(edge[0]) == find(edge[1]) {
			return false
		}
		union(edge[0], edge[1])
	}

	return n == 1
}
