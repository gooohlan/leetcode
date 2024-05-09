package DP

func strStr(haystack string, needle string) int {
    m := len(haystack)
    n := len(needle)
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, 27)
    }
    dp[0][needle[0]-'a'] = 1
    X := 0
    
    for j := 1; j < n; j++ {
        for c := 0; c < 26; c++ {
            dp[j][c] = dp[X][c]
        }
        dp[j][needle[j]-'a'] = j + 1
        X = dp[X][needle[j]-'a']
    }
    
    j := 0
    for i := 0; i < m; i++ {
        j = dp[j][haystack[i]-'a']
        if j == n {
            return i - n + 1
        }
    }
    return -1
}
