package WeeklyContest

import "math"

func paintWalls(cost, time []int) int {
    n := len(cost)
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, 2*n+1) // time[i] <= 500
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    
    var dfs func(int, int) int // 刷前 i 堵墙,且付费时间和为 j 最少开销
    dfs = func(i, j int) int {
        if j-n > i { // 入口处j+n产生了偏移量
            return 0 // 剩余的墙都可以免费刷
        }
        
        if i < 0 {
            return math.MaxInt / 2
        }
        
        if memo[i][j] == -1 {
            memo[i][j] = min(dfs(i-1, j+time[i])+cost[i], dfs(i-1, j-1))
        }
        return memo[i][j]
    }
    return dfs(n-1, n) // 加上偏移量 n，防止负数
}

func paintWallsDP(cost, time []int) int {
    n := len(cost)
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, 2*n+1)
        for j := range dp[i] {
            dp[i][j] = math.MaxInt / 2
        }
    }
    dp[0][n] = 0 // 正常应该是dp[0][0],防止负数,加上偏移量
    for i := 0; i < n; i++ {
        for j := 1; j < 2*n+1; j++ {
            dp[i+1][j-1] = min(dp[i+1][j-1], dp[i][j])
            ne := min(2*n, j+time[i])
            dp[i+1][ne] = min(dp[i+1][ne], dp[i][j]+cost[i])
        }
    }
    res := math.MaxInt
    for j := n; j < 2*n+1; j++ {
        res = min(res, dp[n][j])
    }
    return res
}
