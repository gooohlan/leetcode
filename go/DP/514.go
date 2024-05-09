package DP

import "math"

func findRotateSteps(ring string, key string) int {
    indexMap := make([][]int, 26)
    for i, s := range ring {
        c := s - 'a'
        indexMap[c] = append(indexMap[c], i)
    }
    m, n := len(ring), len(key)
    memo := make([][]int, m)
    for i := range memo {
        memo[i] = make([]int, n)
    }
    
    var dfs func(i, j int) int // 圆盘指针在ring[i], key[j]最小操作数
    dfs = func(i, j int) int {
        if j == n {
            return 0
        }
        
        if memo[i][j] == 0 {
            cur := key[j] - 'a'
            res := math.MaxInt
            for _, target := range indexMap[cur] {
                distance := abs(i - target)
                distance = min(distance, m-distance)        // 顺时针和逆时针选择一个最小的
                res = min(res, 1+distance+dfs(target, j+1)) // 调到ring[target],继续寻找key[k+1]
            }
            memo[i][j] = res
        }
        
        return memo[i][j]
    }
    
    return dfs(0, 0)
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func findRotateStepsDP1(ring string, key string) int {
    indexMap := make([][]int, 26)
    for i, s := range ring {
        c := s - 'a'
        indexMap[c] = append(indexMap[c], i)
    }
    
    m, n := len(ring), len(key)
    dp := make([][]int, m+1) // dp[i][j] 表示当前ring指针指向 ring[i] 需要凑出 key[j..] 需要的最少步骤
    for i := range dp {
        dp[i] = make([]int, n+1)
        for j := 0; j <= n; j++ {
            dp[i][j] = math.MaxInt / 2
        }
    }
    for i := 0; i < m; i++ {
        dp[i][n] = 0
    }
    
    for j := n - 1; j >= 0; j-- {
        for i := 0; i < m; i++ {
            for _, k := range indexMap[key[j]-'a'] {
                dp[i][j] = min(dp[i][j], dp[k][j+1]+min(abs(k-i), m-abs(k-i))+1)
            }
        }
    }
    
    return dp[0][0]
}

func findRotateStepsDP2(ring string, key string) int {
    indexMap := make([][]int, 26)
    for i, s := range ring {
        c := s - 'a'
        indexMap[c] = append(indexMap[c], i)
    }
    
    m, n := len(key), len(ring)
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, n)
        for j := range dp[i] {
            dp[i][j] = math.MaxInt / 2
        }
    }
    
    for _, s := range indexMap[key[0]-'a'] {
        dp[0][s] = min(s, n-s) + 1
    }
    
    for i := 1; i < m; i++ {
        for _, j := range indexMap[key[i]-'a'] {
            for _, k := range indexMap[key[i-1]-'a'] {
                dp[i][j] = min(dp[i][j], dp[i-1][k]+min(abs(j-k), n-abs(j-k))+1)
            }
        }
    }
    
    return min(dp[m-1]...)
}

func findRotateStepsDP2_2(ring string, key string) int {
    indexMap := make([][]int, 26)
    for i, s := range ring {
        c := s - 'a'
        indexMap[c] = append(indexMap[c], i)
    }
    
    m, n := len(key), len(ring)
    dp := make([]int, n)
    for i := range dp {
        dp[i] = math.MaxInt / 2
    }
    
    for _, s := range indexMap[key[0]-'a'] {
        dp[s] = min(s, n-s) + 1
    }
    
    for i := 1; i < m; i++ {
        pre := make([]int, n)
        for i := range pre {
            pre[i] = math.MaxInt / 2
        }
        for _, j := range indexMap[key[i]-'a'] {
            for _, k := range indexMap[key[i-1]-'a'] {
                pre[j] = min(pre[j], dp[k]+min(abs(j-k), n-abs(j-k))+1)
            }
        }
        dp = pre
    }
    
    return min(dp...)
}
