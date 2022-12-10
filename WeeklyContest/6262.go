package WeeklyContest

import "math"

func maxStarSum(vals []int, edges [][]int, k int) int {
	var (
		graph = make([][]int, len(vals))
		// visited = make(map[int]bool) // 记录遍历过的节点，防止走回头路
		onPath = make(map[int]bool) // 记录一次递归堆栈中的节点
		stack  []int                // 记录当前遍历中存储的节点
		max    = math.MinInt
		dfs    func(int)
	)

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	dfs = func(i int) {
		if max < vals[i] {
			max = vals[i]
		}
		if onPath[i] {
			return
		}
		// visited[i] = true
		onPath[i] = true
		stack = append(stack, i)
		if len(stack) > k {
			sum := maxSubArray(vals, stack[len(stack)-k-1:])
			if max < sum {
				max = sum
			}
		}

		for _, cur := range graph[i] {
			// if !visited[cur] {
			dfs(cur)
			// }
		}
		stack = stack[:len(stack)-1]
		onPath[i] = false
	}

	for i := 0; i < len(vals); i++ {
		// if !visited[i] {
		dfs(i)
		// }
	}

	return max
}

func sumStack(vas, stack []int) (sum int) {
	for _, i := range stack {
		sum += vas[i]
	}
	return
}

func maxSubArray(vals, nums []int) int {
	max := vals[nums[0]]
	for i := 1; i < len(nums); i++ {
		if vals[nums[i]]+vals[nums[i-1]] > vals[nums[i]] {
			vals[nums[i]] += vals[nums[i-1]]
		}
		if vals[nums[i]] > max {
			max = vals[nums[i]]
		}
	}
	return max
}
