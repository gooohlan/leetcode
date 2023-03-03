package DFS

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
    var (
        res  int
        m, n = len(grid1), len(grid1[0])
        dfs  func(i, j int)
    )
    
    dfs = func(i, j int) {
        if i < 0 || j < 0 || i >= m || j >= n {
            return
        }
        if grid2[i][j] == 0 {
            return
        }
        
        grid2[i][j] = 0
        for _, dir := range dirs {
            dfs(i+dir[0], j+dir[1])
        }
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid1[i][j] == 0 && grid2[i][j] == 1 {
                dfs(i, j)
            }
        }
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid2[i][j] == 1 {
                res++
                dfs(i, j)
            }
        }
    }
    
    return res
}
