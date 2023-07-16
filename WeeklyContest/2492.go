package WeeklyContest

import "math"

type road struct {
	to, distance int
}

func minScore(n int, roads [][]int) int {
	var (
		graph       = make([][]road, n+1)
		dfs         func(int)
		minDistance = math.MaxInt
		visited     = make(map[int]bool)
	)

	for _, item := range roads {
		graph[item[0]] = append(graph[item[0]], road{item[1], item[2]})
		graph[item[1]] = append(graph[item[1]], road{item[0], item[2]})
	}

	dfs = func(i int) {
		visited[i] = true
		for _, t := range graph[i] {
			minDistance = min(minDistance, t.distance)
			if !visited[t.to] {
				dfs(t.to)
			}
		}
	}

	dfs(1)
	return minDistance
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
