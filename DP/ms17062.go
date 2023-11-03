package DP

import "strconv"

func numberOf2sInRange(n int) int {
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
    f = func(i, cnt int, isLimit bool) (res int) {
        if i == l {
            return cnt
        }
        
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
        
        for d := 0; d <= up; d++ {
            c := cnt
            if d == 2 {
                c++
            }
            res += f(i+1, c, isLimit && d == up)
        }
        return
    }
    return f(0, 0, true)
}
