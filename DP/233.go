package DP

import "strconv"

func countDigitOne(n int) int {
    s := strconv.Itoa(n)
    l := len(s)
    dp := make([][]int, l)
    for i := range dp {
        dp[i] = make([]int, l)
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }
    
    var f func(int, int, bool) int
    f = func(i int, cnt int, isLimit bool) int {
        if i == l {
            return cnt
        }
        var res int
        if !isLimit {
            p := &dp[i][cnt]
            if *p >= 0 {
                return *p
            }
            defer func() { *p = res }()
        }
        
        up := 9
        if isLimit {
            up = int(s[i] - '0')
        }
        for d := 0; d <= up; d++ { // 枚举要填入的数字 d
            c := cnt
            if d == 1 { // 满足要求结果+1
                c++
            }
            res += f(i+1, c, isLimit && d == up)
        }
        return res
    }
    return f(0, 0, true)
}
