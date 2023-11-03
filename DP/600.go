package DP

import "strconv"

func findIntegers(n int) int {
    s := strconv.FormatInt(int64(n), 2)
    l := len(s)
    dp := make([][2]int, l)
    for i := range dp {
        dp[i] = [2]int{-1, -1}
    }
    
    var f func(int, int8, bool) int
    f = func(i int, pre int8, isLimit bool) (res int) {
        if i == l {
            return 1
        }
        
        if !isLimit {
            p := &dp[i][pre]
            if *p >= 0 {
                return *p
            }
            defer func() { *p = res }()
        }
        up := 1
        if isLimit {
            up = int(s[i] & 1)
        }
        res = f(i+1, 0, isLimit && up == 0) // 填入0
        if pre == 0 && up == 1 {            // 前一位是0,并且这一位最高可以填1
            res += f(i+1, 1, isLimit)
        }
        return
    }
    return f(0, 0, true)
}
