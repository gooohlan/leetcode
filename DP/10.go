package DP

import "fmt"

func isMatch(s string, p string) bool {
    m, n := len(s), len(p)
    memo := make(map[string]bool)
    
    var dfs func(i,j int) bool
    dfs = func(i, j int) bool {
        // base case
        if j == n{
            return i == m
        }
        
        if i == m {// 当s走到末尾,p为走到末尾时,p只能剩下a*b*这种格式
            if (n-j)%2 == 1 {
                return false
            }
            for ;j+1 < n; j+=2{
                if p[j+1] != '*'{
                    return false
                }
            }
            return true
        }
        
        key := fmt.Sprintf("%d%d",i,j)
        if _, ok := memo[key]; !ok {
            if s[i] == p[j] || p[j] == '.' { // 字符相等匹配
                if j < n -1 && p[j+1] == '*' {
                    memo[key] = dfs(i, j+2) || dfs(i+1, j) // 通配符匹配 0 次或多次
                }else{
                    memo[key] = dfs(i+1, j+1) // 常规匹配,向后移动
                }
            }else{
                if j < n-1  && p[j+1] == '*'{ // 当前字符不匹配,但正则下个字符是'*'
                    memo[key] =  dfs(i, j+2)
                }else{
                    memo[key] = false
                }
            }
        }
        
        return memo[key]
    }
    
    
    return dfs(0, 0)
}
