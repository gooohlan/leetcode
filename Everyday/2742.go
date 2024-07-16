package Everyday

import "math"

func paintWalls(cost []int, time []int) int {
    n := len(cost)
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, n*2)
        for j := range memo[i] {
            memo[i][j] = -1 // 没有计算过
        }
    }
    
    var dfs func(int, int) int
    dfs = func(i int, j int) int {
        if j > i { // 剩余的墙可以免费刷
            return 0
        }
        if i < 0 { // 表示j<i<0,不符合题目要求
            return math.MaxInt / 2
        }
        
        if memo[i][j+n] == -1 { // 加上n 防止出现负数
            memo[i][j+n] = min(dfs(i-1, j+time[i])+cost[i], dfs(i-1, j-1))
        }
        
        return memo[i][j+n]
    }
    
    return dfs(n-1, 0)
}
