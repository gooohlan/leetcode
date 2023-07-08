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
        if j-n >= n-i { // 入口处j+n产生了偏移量
            return 0 // 剩余的墙都可以免费刷
        }
        
        if i == n {
            return math.MaxInt / 2
        }
        
        if memo[i][j] == -1 {
            memo[i][j] = min(dfs(i+1, j+time[i])+cost[i], dfs(i+1, j-1))
        }
        return memo[i][j]
    }
    return dfs(0, n) // 加上偏移量 n，防止负数
}

func paintWalls1(cost, time []int) int {
    n := len(cost)
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, n+1)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    var dfs func(int, int) int
    dfs = func(i int, j int) int {
        if j <= 0 { // 完成所需选择物品,后面什么也不用选了
            return 0
        }
        if i < 0 { // 已经没有物品可选,完成j所需,不合法
            return math.MaxInt / 2
        }
        
        if memo[i][j] == -1 {
            memo[i][j] = min(dfs(i-1, j-time[i]-1)+cost[i], dfs(i-1, j))
        }
        return memo[i][j]
    }
    
    return dfs(n-1, n)
}
func paintWalls1DP(cost, time []int) int {
    n := len(cost)
    dp := make([]int, n+1)
    for i := 1; i <= n; i++ {
        dp[i] = math.MaxInt / 2
    }
    for i, c := range cost {
        t := time[i] + 1
        for j := n; j > 0; j-- {
            dp[j] = min(dp[max(j-t, 0)]+c, dp[j]) // j-t会小于0
            // memo[i][j] = min(dfs(i-1, j-time[i]-1)+cost[i], dfs(i-1, j))
        }
    }
    return dp[n]
}
