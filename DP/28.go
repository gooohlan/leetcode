package DP

func strStr(haystack string, needle string) int {
    m := len(haystack)
    dp := make([][]int, m)
    for i := range dp {
        dp[i] = make([]int, 27)
    }
    dp[0][haystack[0]-'a'] = 1
    X := 0
    
    for j := 1; j < m; j++ {
        for c := 0; c < 26; c++ {
            dp[j][c] = dp[X][c]
        }
        dp[j][haystack[j]-'a'] = j + 1
        X = dp[X][haystack[j]-'a']
    }
}
