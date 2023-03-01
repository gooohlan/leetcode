package DFS

func numEnclaves(grid [][]int) int {
    var (
        res  = 0
        m, n = len(grid), len(grid[0])
        dfs  = func(i, j int) {}
    )
    dfs = func(i, j int) {
        if i < 0 || j < 0 || i >= m || j >= n {
            return
        }
        if grid[i][j] == 0 {
            return
        }
        
        grid[i][j] = 0
        for _, dir := range dirs {
            dfs(i+dir[0], j+dir[1])
        }
    }
    
    for i := 0; i < m; i++ {
        dfs(i, 0)
        dfs(i, n-1)
    }
    for j := 1; j < n; j++ {
        dfs(0, j)
        dfs(m-1, j)
    }
    
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if grid[i][j] == 1 {
                res++
            }
        }
    }
    
    return res
}
