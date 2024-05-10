package DFS

import "strconv"

func numDistinctIslands(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    var dfs func(i, j, d int) string
    dfs = func(i, j, d int) string {
        if i < 0 || j < 0 || i >= m || j >= n {
            return ""
        }
        if grid[i][j] == 0 {
            return ""
        }
        grid[i][j] = 0
        res := strconv.Itoa(d)
        for i, dir := range dirs {
            res += dfs(i+dir[0], j+dir[1], i+1)
        }
        res += strconv.Itoa(-d)
        return res
    }
    
    var set = make(map[string]bool)
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                s := dfs(i, j, 0)
                if len(s) > 0 {
                    set[s] = true
                }
            }
        }
    }
    
    return len(set)
}
