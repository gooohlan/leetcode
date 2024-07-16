package DP

import "math/bits"

func countArrangement(n int) int {
    memo := make([]int, 1<<n)
    for i := range memo {
        memo[i] = -1 // -1 表示没有计算过
    }
    var dfs func(int) int
    dfs = func(s int) int {
        if s == 1<<n-1 { // 所有元素已选择
            return 1
        }
        
        if memo[s] != -1 {
            return memo[s]
        }
        
        res := 0
        i := bits.OnesCount(uint(s)) + 1 // 有多少个1表示选择的第i个
        for j := 1; j <= n; j++ {
            if s>>(j-1)&1 == 0 && (i%j == 0 || j%i == 0) { // 集合存储的是(0,n-1),此处j需要-1处理
                res += dfs(s | 1<<(j-1))
            }
        }
        memo[s] = res
        return res
    }
    return dfs(0)
}
