package DFS

func maxAreaOfIsland(grid [][]int) int {
    var (
        res  int
        m, n = len(grid), len(grid[0])
        dfs  func(i, j int) int
    )
    
    dfs = func(i, j int) int {
        if i < 0 || j < 0 || i >= m || j >= n {
            return 0
        }
        if grid[i][j] == 0 {
            return 0
        }
        
        grid[i][j] = 0
        
        return 1 + dfs(i-1, j) + dfs(i+1, j) + dfs(i, j-1) + dfs(i, j+1)
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                res = max(res, dfs(i, j))
            }
        }
    }
    
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
