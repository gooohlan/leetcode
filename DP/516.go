package DP

func longestPalindromeSubseq(s string) int {
    var dfs func(int, int) int
    memo := make([][]int, len(s))
    for i := range memo {
        memo[i] = make([]int, len(s))
    }
    dfs = func(i, j int) int {
        if i > j {
            return 0
        }
        
        if i == j {
            return 1
        }
        
        if memo[i][j] == 0 {
            if s[i] == s[j] {
                memo[i][j] = dfs(i+1, j-1) + 2
            } else {
                memo[i][j] = max(dfs(i+1, j), dfs(i, j-1))
            }
        }
        return memo[i][j]
    }
    return dfs(0, len(s)-1)
}

func longestPalindromeSubseqDP(s string) int {
    n := len(s)
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
        dp[i][i] = 1
    }
    
    for i := n - 1; i >= 0; i-- {
        for j := i + 1; j < n; j++ { // 只遍历dp[i][i]右侧部分
            if s[i] == s[j] {
                dp[i][j] = dp[i+1][j-1] + 2
            } else {
                dp[i][j] = max(dp[i+1][j], dp[i][j-1])
            }
        }
    }
    
    return dp[0][n-1]
}
