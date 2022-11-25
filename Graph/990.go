package Graph

// 通过'=='建立连通性,再通过'!='判断是否与连通性冲突
func equationsPossible(equations []string) bool {
	var (
		parent = make([]int, 26) // 父节点
		find   func(int) int
		union  func(int, int)
	)

	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union = func(p int, q int) {
		rootP, rootQ := find(p), find(q)
		if rootP == rootQ {
			return
		}

		parent[rootQ] = rootP
	}

	for i := 0; i < 26; i++ {
		parent[i] = i
	}

	for _, str := range equations {
		if str[1] == '=' {
			union(int(str[0]-'a'), int(str[3]-'a'))
		}
	}

	for _, str := range equations {
		if str[1] == '!' {
			if find(int(str[0]-'a')) == find(int(str[3]-'a')) {
				return false
			}
		}
	}
	return true
}
