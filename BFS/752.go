package BFS

func openLock(deadends []string, target string) int {
	queue := []string{"0000"}
	dead := map[string]bool{}
	depth := 0
	// target

	for len(queue) != 0 {
		sz := len(queue)
		for i := 0; i < sz; i++ {
			cur := queue[0]
			queue = queue[1:]
			// 判断是否到达终点
			if cur == target {
				return depth
			}
		}
	}
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

func minusOne(s string, i int) string {
	b := []byte(s)
	b[i] -= 1
	if b[i] < '0' {
		b[i] = '9'
	}
	return string(b)
}
