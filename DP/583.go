package DP

func minDistance583(word1 string, word2 string) int {
    var dfs func(int, int) int
    memo := make([][]int, len(word1))
    for i := 0; i < len(word1); i++ {
        memo[i] = make([]int, len(word2))
        for j := 0; j < len(word2); j++ {
            memo[i][j] = -1
        }
    }
    
    dfs = func(i int, j int) int {
        if i == len(word1) {
            return len(word2) - j
        }
        if j == len(word2) {
            return len(word1) - i
        }
        if memo[i][j] == -1 {
            if word1[i] == word2[j] {
                memo[i][j] = dfs(i+1, j+1)
            } else {
                memo[i][j] = min(dfs(i+1, j), dfs(i, j+1)) + 1
            }
        }
        return memo[i][j]
    }
    
    return dfs(0, 0)
}

func minDistance583DP(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    dp := make([][]int, m+1)
    
    // 初始化dp数组,当i或j为0时,删除未到末尾的数量即可
    for i := range dp {
        dp[i] = make([]int, n+1)
        dp[i][0] = i
    }
    for j := range dp[0] {
        dp[0][j] = j
    }
    
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1
            }
        }
    }
    return dp[m][n]
}
