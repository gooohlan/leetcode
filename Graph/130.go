package Graph

// 将边界上相邻节点与一个父节点dummy相连, 然后遍历所有元素,与dummy不想通的则变为X
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	var (
		parent []int // 父节点
		find   func(int) int
		union  func(int, int)
		xy     = [4][2]int{[2]int{1, 0}, [2]int{-1, 0}, [2]int{0, 1}, [2]int{0, -1}}
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

	n, m := len(board), len(board[0])
	dummy := m * n
	parent = make([]int, m*n+1)
	for i := 0; i < m*n+1; i++ {
		parent[i] = i
	}

	// 将首列和末列的 O 与 dummy 连通
	for i := 0; i < n; i++ { // 标记第一列和最后一列
		if board[i][0] == 'O' {
			union(i*m, dummy)
		}
		if board[i][m-1] == 'O' {
			union(i*m+m-1, dummy)
		}
	}

	// 将首行和末行的 O 与 dummy 连通
	for i := 1; i < m-1; i++ { //
		if board[0][i] == 'O' {
			union(i, dummy)
		}
		if board[n-1][i] == 'O' {
			union(m*(n-1)+i, dummy)
		}
	}

	// 将内部'O'相连,如果内部'O'于边界处相连,就会连接到dummy
	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			if board[i][j] == 'O' {
				for k := 0; k < 4; k++ {
					x, y := i+xy[k][0], j+xy[k][1]
					if board[x][y] == 'O' {
						union(x*m+y, i*m+j)
					}
				}
			}
		}
	}

	// 有不和 dummy 连通的 O，都要被替换
	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			if find(dummy) != find(i*m+j) {
				board[i][j] = 'X'
			}
		}

	}
}

// 先标记与边界上相关的'O',然后再遍历所有的元素,未被标记的元素变为'X',已被标记的不变
func solveDFS(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	var dfs func(int, int)
	n, m := len(board), len(board[0])

	dfs = func(x, y int) {
		if x < 0 || x >= n || y < 0 || y >= m || board[x][y] != 'O' {
			return
		}
		board[x][y] = '#' // 先替换边界处'O'
		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
	}

	for i := 0; i < n; i++ { // 标记第一列和最后一列
		dfs(i, 0)
		dfs(i, m-1)
	}
	for i := 1; i < m-1; i++ { //  标记第一行和最后一行
		dfs(0, i)
		dfs(n-1, i)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'O' { // 未被标记,变成X
				board[i][j] = 'X'
			}
			if board[i][j] == '#' { // 被标记无需改变,变回'O'
				board[i][j] = 'O'
			}
		}
	}
}

func solveBFS(board [][]byte) {
	var ( // id,y分别代表当前元素的4个方向
		xy = [4][2]int{[2]int{1, 0}, [2]int{-1, 0}, [2]int{0, 1}, [2]int{0, -1}}
	)
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	n, m := len(board), len(board[0])
	q := [][]int{}
	// 先标记边界上的'O'
	for i := 0; i < n; i++ { // 标记第一列和最后一列
		if board[i][0] == 'O' {
			q = append(q, []int{i, 0})
			board[i][0] = '#'
		}
		if board[i][m-1] == 'O' {
			q = append(q, []int{i, m - 1})
			board[i][m-1] = '#'
		}
	}
	for i := 1; i < m-1; i++ { //  标记第一行和最后一行
		if board[0][i] == 'O' {
			q = append(q, []int{0, i})
			board[0][i] = '#'
		}
		if board[n-1][i] == 'O' {
			q = append(q, []int{n - 1, i})
			board[n-1][i] = '#'
		}
	}

	// 标记与边界相邻的
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		for i := 0; i < 4; i++ {
			x, y := cur[0]+xy[i][0], cur[1]+xy[i][1]
			if x < 0 || x >= n || y < 0 || y >= m || board[x][y] != 'O' {
				continue
			}
			q = append(q, []int{x, y})
			board[x][y] = '#'
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'O' { // 未被标记,变成X
				board[i][j] = 'X'
			}
			if board[i][j] == '#' { // 被标记无需改变,变回'O'
				board[i][j] = 'O'
			}
		}
	}
}
