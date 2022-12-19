package DP

func minDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    memo := make([][]int, m)
    for i := range memo {
        memo[i] = make([]int, n)
        for j := 0; j < n; j++ {
            memo[i][j] = -1
        }
    }
    
    var dp func(int, int) int
    
    dp = func(i, j int) int {
        if i == -1 {
            return j + 1
        }
        if j == -1 {
            return i + 1
        }
        
        if memo[i][j] != -1 {
            return memo[i][j]
        }
        
        if word1[i] == word2[j] {
            memo[i][j] = dp(i-1, j-1)
        } else {
            memo[i][j] = min(
                dp(i, j-1)+1,
                dp(i-1, j)+1,
                dp(i-1, j-1)+1)
        }
        return memo[i][j]
    }
    return dp(m-1, n-1)
}

func min(arr ...int) int {
    res := arr[0]
    for _, v := range arr {
        if v < res {
            res = v
        }
    }
    return res
}

func minDistanceDP(word1, word2 string) int {
    m, n := len(word1), len(word2)
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    
    for i := 1; i <= m; i++ {
        dp[i][0] = i
    }
    for i := 0; i <= n; i++ {
        dp[0][i] = i
    }
    
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if word1[i-1] == word2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
            }
        }
    }
    return dp[m][n]
}
