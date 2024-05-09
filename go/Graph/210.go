package Graph

// 此题与207几乎一样
func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		// 有向图
		graph    = make([][]int, numCourses)
		dfs      func(int)
		visited  = make(map[int]bool) // 记录遍历过的节点，防止走回头路
		onPath   = make(map[int]bool) // 记录一次递归堆栈中的节点
		hasCycle = false              // 记录图中是否有环
		res      []int
	)

	for _, item := range prerequisites {
		graph[item[1]] = append(graph[item[1]], item[0])
	}

	dfs = func(s int) {
		if onPath[s] {
			hasCycle = true // 出现环
		}

		if visited[s] || hasCycle {
			return // 出现环或者已经找过了
		}

		visited[s] = true
		onPath[s] = true
		for _, t := range graph[s] {
			dfs(t)
		}
		onPath[s] = false
		res = append(res, s)
	}

	for i := 0; i < numCourses && !hasCycle; i++ {
		if !visited[i] {

			dfs(i)
		}
	}
	if hasCycle {
		return []int{}
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}

func findOrderBFS(numCourses int, prerequisites [][]int) []int {
	var (
		graph    = make([][]int, numCourses) // 领接表
		indegree = make([]int, numCourses)   // 入度数
		queue    []int                       // 队列
		res      []int                       // 记录过的节点
	)

	for _, item := range prerequisites {
		graph[item[1]] = append(graph[item[1]], item[0]) // 创建领接表
		indegree[item[0]]++                              // 记录节点入度数
	}
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i) // 初始化队列
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		res = append(res, cur)
		for _, next := range graph[cur] {
			indegree[next]--
			if indegree[next] == 0 { // 如果入度变为 0，说明 next 依赖的节点都已被遍历
				queue = append(queue, next)
			}
		}
	}

	if len(res) != numCourses {
		return []int{}
	}
	return res
}
