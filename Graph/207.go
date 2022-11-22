package Graph

// 根据题意,存在循环依赖则无法上完所有课程
// 将课程转换为领接表,判断是否成环即可
func canFinishDFS(numCourses int, prerequisites [][]int) bool {
	var (
		// 有向图
		graph    = make([][]int, numCourses)
		dfs      func(int)
		visited  = make(map[int]bool) // 记录遍历过的节点，防止走回头路
		onPath   = make(map[int]bool) // 记录一次递归堆栈中的节点
		hasCycle = false              // 记录图中是否有环
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
	}

	for i := 0; i < numCourses; i++ {
		dfs(i)
	}
	return !hasCycle
}

// 利用有向图中的「出度」和「入度」概念
// 还是先构建领接表,并同时构建一个数组记录每个节点的「入度」
// 对BFS的队列进行初始化,将入度为0的额节点装入队列
// 遍历队列开始BFS循环,减少相邻节点的入度,入度为0的节点加入队列
// 所有节点都别遍历过则表示不存在环,反之存在
func canFinishBFS(numCourses int, prerequisites [][]int) bool {
	var (
		graph    = make([][]int, numCourses) // 领接表
		indegree = make([]int, numCourses)   // 入度数
		queue    []int                       // 队列
		count    int                         // 记录过的节点数
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
		count++
		for _, next := range graph[cur] {
			indegree[next]--
			if indegree[next] == 0 { // 如果入度变为 0，说明 next 依赖的节点都已被遍历
				queue = append(queue, next)
			}
		}
	}

	return count == numCourses
}
