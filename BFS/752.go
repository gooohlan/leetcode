package BFS

// 四个数字可以向上或者向下转,这样就形成了一颗8叉树
// 后续的操作无非就是在这颗树上遍历,进行一些限制条件的剪枝之类的操作
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

// 双向BFS
func openLock2(deadends []string, target string) int {
	q1 := map[string]bool{"0000": true}
	q2 := map[string]bool{target: true}

	dead := map[string]bool{}
	for _, s := range deadends {
		dead[s] = true
	}
	if dead[target] {
		return -1
	}

	depth := 0
	seen := map[string]bool{"0000": true}

	for len(q1) != 0 && len(q2) != 0 {
		if len(q1) > len(q2) {
			q1, q2 = q2, q1
		}
		temp := make(map[string]bool)

		for s, _ := range q1 {
			if dead[s] {
				continue
			}

			// 判断是否到达终点
			if q2[s] {
				return depth
			}
			// 已经进行过判断,所以加入已经转到过队列
			seen[s] = true

			// q1与q2需要判断是否相同,所以允许重复
			for i := 0; i < 4; i++ {
				up := plusOne(s, i)
				if !seen[up] {
					temp[up] = true
				}

				down := minusOne(s, i)
				if !seen[down] {
					temp[down] = true
				}
			}

		}

		depth++
		q1, q2 = q2, temp
	}
	return -1
}
