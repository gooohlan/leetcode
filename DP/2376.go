package DP

import "strconv"

func countSpecialNumbers(n int) int {
    s := strconv.Itoa(n)
    l := len(s)
    memo := make([][1 << 10]int, l)
    for i := range memo {
        for j := range memo[i] {
            memo[i][j] = -1
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
            p := &memo[i][mark]
            if *p >= 0 {
                return *p
            }
            defer func() { *p = res }()
        }
        if !isNum { // 跳过当前数位
            res += f(i+1, mark, false, false)
        }
        
        d := 0
        if !isNum {
            d = 1 // 前面没数字,最低填1
        }
        up := 9
        if isLimit {
            up = int(s[i] - '0')
        }
        for ; d <= up; d++ {
            if mark>>d&1 == 0 { // d没选中过
                res += f(i+1, mark|1<<d, isLimit && d == up, true)
            }
        }
        return
    }
    return f(0, 0, true, false)
}
