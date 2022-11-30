package Graph

func findCelebrity(n int) int {
	res := 0 // 假设res是名人
	for i := 1; i < n; i++ {
		if knows(res, i) || !knows(i, res) { // 如果res认识i 或者i不会认识res,那么res肯定不是名人
			res = i
		}
	}

	for i := 0; i < n; i++ {
		if i == res {
			continue
		}
		if knows(res, i) || !knows(i, res) {
			return -1
		}
	}
	return res
}

var graph = [][]int{
	[]int{1, 1, 1, 0},
	[]int{1, 1, 1, 1},
	[]int{0, 0, 1, 0},
	[]int{0, 0, 1, 1},
}

func knows(a, b int) bool {
	if graph[a][b] == 0 {
		return false
	}
	return true
}
