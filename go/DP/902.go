package DP

import "strconv"

func atMostNGivenDigitSet(digits []string, n int) int {
    s := strconv.Itoa(n)
    l := len(s)
    dp := make([]int, l)
    for i := range dp {
        dp[i] = -1
    }
    
    var f func(int, bool, bool) int
    f = func(i int, isLimit bool, isNum bool) (res int) {
        if i == l {
            if isNum {
                return 1
            }
            return
        }
        if !isLimit && isNum { // 不受到任何约束
            p := &dp[i]
            if *p >= 0 {
                return *p
            }
            defer func() { *p = res }()
        }
        
        if !isNum {
            // isLimit 因为当前位置没有填数字,可以跳过
            res += f(i+1, false, false)
        }
        // 根据是否收到约束,决定可以填入数字的上线
        up := byte('9')
        if isLimit {
            up = s[i]
        }
        
        // 对于一般的题目而言,如果此时 isNum 为 false, 则必须从1开始枚举,本体中没有0,所以无需处理这种情况
        for _, d := range digits {
            if d[0] > up {
                break
            }
            res += f(i+1, isLimit && up == d[0], true)
        }
        return
    }
    return f(0, true, false)
}
