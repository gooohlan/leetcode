package DP

import "math"

func superEggDrop(k int, n int) int {
    var dfs func(int, int) int
    memo := make([][]int, k+1)
    for i := range memo {
        memo[i] = make([]int, n+1)
    }
    dfs = func(K int, N int) int {
        if K == 1 {
            return N
        }
        if N == 0 {
            return 0
        }
        if memo[K][N] == 0 {
            res := math.MaxInt
            
            for i := 1; i <= N; i++ {
                res = min(res, max(
                    dfs(K, N-i),   // 没碎
                    dfs(K-1, i-1), // 碎了
                )+1)
            }
            memo[K][N] = res
        }
        
        return memo[K][N]
    }
    
    return dfs(k, n)
}
