package DP

func longestCommonSubsequence(text1 string, text2 string) int {
    var dp func(int, int) int
    memo := make([][]int, len(text1))
    for i := 0; i < len(text1); i++ {
        memo[i] = make([]int, len(text2))
        for j := 0; j < len(text2); j++ {
            memo[i][j] = -1
        }
    }
    
    dp = func(i int, j int) int {
        if i == len(text1) || j == len(text2) {
            return 0
        }
        if memo[i][j] == -1 {
            if text1[i] == text2[j] {
                memo[i][j] = 1 + dp(i+1, j+1)
            } else {
                memo[i][j] = max(
                    dp(i+1, j),
                    dp(i, j+1),
                    // dp(i+1, j+1)
                )
            }
        }
        return memo[i][j]
    }
    
    return dp(0, 0)
}

func max(arr ...int) int {
    res := arr[0]
    for _, v := range arr {
        if v > res {
            res = v
        }
    }
    return res
}

func longestCommonSubsequenceDP(text1 string, text2 string) int {
    m, n := len(text1), len(text2)
    dp := make([][]int, m+1)
    for i := 0; i <= m; i++ {
        dp[i] = make([]int, n+1)
    }
    
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if text1[i-1] == text2[j-1] {
                dp[i][j] = 1 + dp[i-1][j-1]
            } else {
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }
    
    return dp[m][n]
}
