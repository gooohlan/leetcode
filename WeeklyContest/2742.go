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
    
    var dfs func(int, int) int // 刷签 i 堵墙,且付费时间和为 j 最少开销
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
