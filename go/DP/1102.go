package DP

import "strconv"

func numDupDigitsAtMostN(n int) int {
    s := strconv.Itoa(n)
    l := len(s)
    dp := make([][1 << 10]int, l)
    for i := range dp {
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }
    var f func(int, int, bool, bool) int
    f = func(i int, mark int, isLimit bool, isNum bool) (res int) {
        if i == l {
            if isNum {
                return 1
            }
            return
        }
        
        if !isLimit && isNum {
            p := &dp[i][mark]
            if *p >= 0 {
                return *p
            }
            defer func() { *p = res }()
        }
        
        if !isNum {
            res += f(i+1, mark, false, false)
        }
        d := 0
        if !isNum {
            d = 1
        }
        up := 9
        if isLimit {
            up = int(s[i] - '0')
        }
        for ; d <= up; d++ {
            if mark>>d&1 == 0 { // d 不在mark中
                res += f(i+1, mark|1<<d, isLimit && d == up, true)
            }
        }
        return
    }
    // 求出所有不重复的数字,减去即可
    return n - f(0, 0, true, false)
}
