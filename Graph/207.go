package Graph

// 根据题意,存在循环依赖则无法上完所有课程
// 将课程转换为有向图,判断是否成环即可
func canFinish(numCourses int, prerequisites [][]int) bool {
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
