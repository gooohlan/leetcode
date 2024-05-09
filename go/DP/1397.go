package DP

import "strings"

const mod = 1e9 + 7

func findGoodStrings(n int, s1 string, s2 string, evil string) int {
    m := len(evil)
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, m)
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }
    
    next := make([]int, m)
    for i := 1; i < m; i++ {
        j := next[i-1]
        for j > 0 && evil[i] != evil[j] {
            j = next[j-1]
        }
        if evil[i] == evil[j] {
            j++
        }
        next[i] = j
    }
    
    var f func(int, int, bool, string) int
    f = func(i int, pos int, isLimit bool, s string) (res int) {
        if pos == m {
            return
        }
        if i == n {
            return 1
        }
        
        if !isLimit {
            p := &dp[i][pos]
            if *p >= 0 {
                return *p
            }
            defer func() { *p = res }()
        }
        
        d := byte(97)
        up := byte(122)
        if isLimit {
            up = s[i]
        }
        for ; d <= up; d++ {
            nxt := pos
            for nxt > 0 && d != evil[nxt] {
                nxt = next[nxt-1]
            }
            // 此处要注意，当 nxt == 0 的时候，会存在 d != evil[nxt] 的情况
            // 若直接 nxt + 1 进入递归，是认为此时的两个字符一定是匹配上了，实际上可能并没有
            if nxt == 0 && d != evil[nxt] {
                nxt = -1
            }
            res += f(i+1, nxt+1, isLimit && d == up, s)
            res %= mod
        }
        return
    }
    res := f(0, 0, true, s2) - f(0, 0, true, s1)
    if res < 0 {
        res += mod
    }
    if strings.Index(s1, evil) == -1 {
        res += 1
    }
    return res
}
