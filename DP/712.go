package DP

func minimumDeleteSum(s1 string, s2 string) int {
    var dp func(int, int) int
    memo := make([][]int, len(s1))
    for i := 0; i < len(s1); i++ {
        memo[i] = make([]int, len(s2))
        for j := 0; j < len(s2); j++ {
            memo[i][j] = -1
        }
    }
    
    dp = func(i int, j int) int {
        var res int
        if i == len(s1) {
            for ; j < len(s2); j++ {
                res += int(s2[j])
            }
            return res
        }
        if j == len(s2) {
            for ; i < len(s1); i++ {
                res += int(s1[i])
            }
            return res
        }
        if memo[i][j] == -1 {
            if s1[i] == s2[j] {
                memo[i][j] = dp(i+1, j+1)
            } else {
                memo[i][j] = min(
                    int(s1[i])+dp(i+1, j),
                    int(s2[j])+dp(i, j+1),
                )
            }
        }
        
        return memo[i][j]
    }
    
    return dp(0, 0)
}

func minimumDeleteSumDP(s1, s2 string) int {
    m, n := len(s1), len(s2)
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
        if i > 0 {
            dp[i][0] = dp[i-1][0] + int(s1[i-1])
        }
    }
    for j := range dp[0] {
        if j > 0 {
            dp[0][j] = dp[0][j-1] + int(s2[j-1])
        }
    }
    
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if s1[i-1] == s2[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = min(
                    int(s1[i-1])+dp[i-1][j],
                    int(s2[j-1])+dp[i][j-1],
                )
            }
        }
    }
    return dp[m][n]
}
