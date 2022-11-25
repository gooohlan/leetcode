package Graph

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

	for i := 0; i < n; i++ { // 标记第一行和最后一行
		dfs(i, 0)
		dfs(i, m-1)
	}
	for i := 1; i < m-1; i++ { //  标记第一列和最后一列
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
