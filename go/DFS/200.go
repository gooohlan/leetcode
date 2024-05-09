package DFS

var dirs = [][]int{
    {-1, 0},
    {1, 0},
    {0, -1},
    {0, 1},
}

func numIslands(grid [][]byte) int {
    res := 0
    m, n := len(grid), len(grid[0])
    var dfs func(i, j int)
    dfs = func(i, j int) {
        if i < 0 || j < 0 || i >= m || j >= n {
            return
        }
        
        if grid[i][j] == '0' {
            return
        }
        
        grid[i][j] = '0'
        for _, dir := range dirs {
            dfs(i+dir[0], j+dir[1])
        }
    }
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == '1' {
                res++
                dfs(i, j)
            }
        }
    }
    
    return res
}
