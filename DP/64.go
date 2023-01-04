package DP

import "math"

func minPathSum(grid [][]int) int {
    var dp func(int, int) int
    m, n := len(grid), len(grid[0])
    memo := make([][]int, m)
    for i := range memo {
        memo[i] = make([]int, n)
    }
    
    dp = func(i int, j int) int {
        if i == m-1 && j == n-1 {
            return grid[m-1][n-1]
        }
        if i == m || j == n {
            return math.MaxInt
        }
        if memo[i][j] == 0 {
            memo[i][j] = grid[i][j] + min(dp(i, j+1), dp(i+1, j))
        }
        
        return memo[i][j]
    }
    
    return dp(0, 0)
}

func minPathSumDP(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    
    dp[0][0] = grid[0][0]
    
    for i := 1; i < m; i++ {
        dp[i][0] = dp[i-1][0] + grid[i][0]
    }
    
    for i := 1; i < n; i++ {
        dp[0][i] = dp[0][i-1] + grid[0][i]
    }
    
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
        }
    }
    
    return dp[m-1][n-1]
}
