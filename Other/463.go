package Other

// 公式 4*岛数量-2*相邻边
func islandPerimeter(grid [][]int) int {
    var count, edge int
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == 0 {
                continue
            }
            count++
            if j+1 < len(grid[i]) && grid[i][j+1] == 1 {
                edge++
            }
            if i+1 < len(grid) && grid[i+1][j] == 1 {
                edge++
            }
        }
    }
    return 4*count - 2*edge
}

func islandPerimeter2(grid [][]int) int {
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == 1 {
                return dfs(grid, i, j)
            }
        }
    }
    return 0
}

func dfs(grid [][]int, x, y int) int {
    if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
        return 1
    }
    if grid[x][y] == 0 {
        return 1
    }
    if grid[x][y] == 2 {
        return 0
    }
    grid[x][y] = 2
    return dfs(grid, x-1, y) + dfs(grid, x+1, y) + dfs(grid, x, y-1) + dfs(grid, x, y+1)
}
