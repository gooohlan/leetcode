package DP

import "math"

func calculateMinimumHP(dungeon [][]int) int {
    var dfs func(i, j int) int // 从i,j到右下角最低消耗血量
    m, n := len(dungeon), len(dungeon[0])
    memo := make([][]int, m)
    for i := range memo {
        memo[i] = make([]int, n)
        for j := 0; j < n; j++ {
            memo[i][j] = -1
        }
    }
    
    dfs = func(i, j int) int {
        if i == m || j == n {
            return math.MaxInt
        }
        
        if i == m-1 && j == n-1 {
            if dungeon[i][j] <= 0 {
                return -dungeon[i][j]
            }
            return 0
        }
        
        if memo[i][j] == -1 {
            res := min(dfs(i+1, j), dfs(i, j+1)) - dungeon[i][j]
            
            if res < 0 {
                res = 0
            }
            memo[i][j] = res
        }
        
        return memo[i][j]
    }
    
    return dfs(0, 0) + 1
}

func calculateMinimumHPDP(dungeon [][]int) int {
    m, n := len(dungeon), len(dungeon[0])
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    dp[m-1][n-1] = max(0, -dungeon[m-1][n-1])
    
    for i := m - 2; i >= 0; i-- {
        dp[i][n-1] = max(0, dp[i+1][n-1]-dungeon[i][n-1])
    }
    
    for i := n - 2; i >= 0; i-- {
        dp[m-1][i] = max(0, dp[m-1][i+1]-dungeon[m-1][i])
    }
    
    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            dp[i][j] = max(0, min(dp[i+1][j], dp[i][j+1])-dungeon[i][j])
        }
    }
    
    return dp[0][0] + 1
}
