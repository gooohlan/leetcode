package BFS

func openLock(deadends []string, target string) int {
	queue := []string{"0000"}

	dead := map[string]bool{}
	for _, s := range deadends {
		dead[s] = true
	}
	if dead[target] {
		return -1
	}
	depth := 0
	seen := map[string]bool{queue[0]: true}

	for len(queue) != 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			cur := queue[0]
			queue = queue[1:]
			// 判断是否到达终点
			if cur == target {
				return depth
			}
			// 死锁节点跳过
			if dead[cur] {
				continue
			}

			for i := 0; i < 4; i++ {
				up := plusOne(cur, i)
				// 为转到过则添加到队列
				if !seen[up] {
					queue = append(queue, up)
					seen[up] = true
				}

				down := minusOne(cur, i)
				if !seen[down] {
					queue = append(queue, down)
					seen[down] = true
				}
			}
		}
		depth++
	}
	return -1
}

// 将s[i]向上拨
func plusOne(s string, i int) string {
	b := []byte(s)
	b[i] += 1
	if b[i] > '9' {
		b[i] = '0'
	}
	return string(b)
}

// 将s[i]向下拨
func minusOne(s string, i int) string {
	b := []byte(s)
	b[i] -= 1
	if b[i] < '0' {
		b[i] = '9'
	}
	return string(b)
}
