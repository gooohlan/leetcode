package DP

import "fmt"

func isMatch(s string, p string) bool {
    m, n := len(s), len(p)
    memo := make(map[string]bool)
    
    var dfs func(i, j int) bool
    dfs = func(i, j int) bool {
        // base case
        if j == n {
            return i == m
        }
        
        if i == m { // 当s走到末尾,p为走到末尾时,p只能剩下a*b*这种格式
            if (n-j)%2 == 1 {
                return false
            }
            for ; j+1 < n; j += 2 {
                if p[j+1] != '*' {
                    return false
                }
            }
            return true
        }
        
        key := fmt.Sprintf("%d%d", i, j)
        if _, ok := memo[key]; !ok {
            if s[i] == p[j] || p[j] == '.' { // 字符相等匹配
                if j < n-1 && p[j+1] == '*' {
                    memo[key] = dfs(i, j+2) || dfs(i+1, j) // 通配符匹配 0 次或多次
                } else {
                    memo[key] = dfs(i+1, j+1) // 常规匹配,向后移动
                }
            } else {
                if j < n-1 && p[j+1] == '*' { // 当前字符不匹配,但正则下个字符是'*'
                    memo[key] = dfs(i, j+2)
                } else {
                    memo[key] = false
                }
            }
        }
        
        return memo[key]
    }
    
    return dfs(0, 0)
}

func isMatchDP(s string, p string) bool {
    m, n := len(s), len(p)
    dp := make([][]bool, m+1)
    for i := range dp {
        dp[i] = make([]bool, n+1)
    }
    
    dp[0][0] = true
    
    // 当s为空时,p右端为"*"则可以匹配
    for j := 1; j <= n; j++ {
        if p[j-1] == '*' {
            dp[0][j] = dp[0][j-2]
        }
    }
    
    for i := 1; i <= m; i++ {
        for j := 1; j <= n; j++ {
            if s[i-1] == p[j-1] || p[j-1] == '.' { // 字段匹配,直接通过
                dp[i][j] = dp[i-1][j-1]
            } else if p[j-1] == '*' { // 当前字符匹配,则查看p当前是否为*
                if s[i-1] == p[j-2] || p[j-2] == '.' { // 看星号前一个字符是否匹配
                    // dp[i][j-2] 表示*只重复0次
                    // dp[i-1][j-2] 重复一次
                    // dp[i-1][j] 重复次数大于1
                    dp[i][j] = dp[i][j-2] || dp[i-1][j-2] || dp[i-1][j]
                }else{ // 星号前一个字符不匹配,则当星号展示0次处理
                    dp[i][j] = dp[i][j-2]
                }
            }
        }
    }
    return dp[m][n]
}


func isMatchDP2(s string, p string) bool {
    m, n := len(s), len(p)
    dp := make([]bool, n+1)
    // for i := range dp {
    //     dp[i] = make([]bool, n+1)
    // }
    
    dp[0] = true
    
    // 当s为空时,p右端为"*"则可以匹配
    for j := 1; j <= n; j++ {
        if p[j-1] == '*' {
            dp[j] = dp[j-2]
        }
    }
    
    for i := 1; i <= m; i++ {
        pre := make([]bool, n+1)
        for j := 1; j <= n; j++ {
            if s[i-1] == p[j-1] || p[j-1] == '.' { // 字段匹配,直接通过
                // dp[i][j] = dp[i-1][j-1]
                pre[j] = dp[j-1]
            } else if p[j-1] == '*' { // 当前字符匹配,则查看p当前是否为*
                if s[i-1] == p[j-2] || p[j-2] == '.' { // 看星号前一个字符是否匹配
                    // dp[i][j-2] 表示*只重复0次
                    // dp[i-1][j-2] 重复一次
                    // dp[i-1][j] 重复次数大于1
                    // dp[i][j] = dp[i][j-2] || dp[i-1][j-2] || dp[i-1][j]
                    pre[j] = pre[j-2] || dp[j-2] || dp[j]
                }else{ // 星号前一个字符不匹配,则当星号展示0次处理
                    pre[j] = pre[j-2]
                }
            }
        }
        dp = pre
    }
    return dp[n]
}
