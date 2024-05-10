package DP

// 这题与516机会一模一样
// 判断s[i]与s[j]是否相等,不相等时
// 在s[j]后面插入s[i],次数为dfs(i+1, j)+1
// 在s[i]前面插入s[j],次数为dfs(i, j-1)+1
// 二者求最小就是答案
func minInsertions(s string) int {
    n := len(s)
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    for i := n - 1; i >= 0; i-- {
        for j := i + 1; j < n; j++ { // 只遍历dp[i][i]右侧部分
            if s[i] == s[j] {
                dp[i][j] = dp[i+1][j-1]
            } else {
                dp[i][j] = min(dp[i+1][j], dp[i][j-1]) + 1
            }
        }
    }
    
    return dp[0][n-1]
}

// 另外一个思路,字符串长度减去最长最长回文子序列长度就是答案
// 因为只需要在非最长子序列位置插入对应字符即可
func minInsertions2(s string) int {
    return len(s) - longestPalindromeSubseqDP(s)
}
