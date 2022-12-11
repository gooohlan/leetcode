package WeeklyContest

func maxPoints(grid [][]int, queries []int) []int {
	m := len(grid)
	n := len(grid[0])
	res := make([]int, len(queries))
	for i, querie := range queries {
		if grid[0][0] >= querie {
			continue
		}
		q := []int{0}
		onPath := make(map[int]bool)
		onPath[0] = true
		res[i]++
		for len(q) != 0 {
			cur := q[0]
			q = q[1:]
			x, y := cur/n, cur%n

			if x+1 < m && grid[x+1][y] < querie && !onPath[cur+n] {
				q = append(q, cur+n)
				res[i]++
				onPath[cur+n] = true
			}

			if x-1 > 0 && grid[x-1][y] < querie && !onPath[cur-n] {
				q = append(q, cur-n)
				res[i]++
				onPath[cur-n] = true
			}

			if y+1 < n && grid[x][y+1] < querie && !onPath[cur+1] {
				q = append(q, cur+1)
				res[i]++
				onPath[cur+1] = true
			}

			if y-1 > 0 && grid[x][y-1] < querie && !onPath[cur-1] {
				q = append(q, cur-1)
				res[i]++
				onPath[cur-1] = true
			}

		}
	}
	return res
}
