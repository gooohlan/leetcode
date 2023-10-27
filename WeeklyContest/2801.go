package WeeklyContest

// const mod = 1e9 + 7

func countSteppingNumbers(low string, high string) int {
    res := (calc(high) - calc(low) + mod) % mod
    if valid(low) {
        res = (res + 1) % mod
    }
    return res
}

func calc(s string) int {
    memo := make([][10]int, len(s))
    for i := range memo {
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    
    // pre 表示上一个数位的值, isNumber为false,可以忽略
    // isLimit 表示当前是否收到了前n位的约束
    // isNumber 表示i前面的数值是否填充了数字
    var dfs func(i, pre int, isLimit, isNumber bool) int
    dfs = func(i, pre int, isLimit, isNumber bool) int {
        if i == len(s) { // 递归结束
            if isNumber {
                return 1
            }
            return 0 // 前面最后一位没填充数字,非法
        }
        
        var res int
        if !isLimit && isNumber {
            p := &memo[i][pre]
            if *p >= 0 {
                return *p
            }
            defer func() { *p = res }()
        }
        
        if !isNumber { // 这一位也不填,直接跳过
            res += dfs(i+1, pre, false, false)
        }
        
        d := 0
        if !isNumber {
            d = 1 // 前面没填,这一位必须从1开始
        }
        up := 9
        if isLimit {
            up = int(s[i] - '0') // 前面选择的字符与 s 一样,这意味最多只能填 s[i]
        }
        for ; d <= up; d++ {
            if !isNumber || abs(d, pre) == 1 { // 第一位数字随便填，其余必须相差 1
                res += dfs(i+1, d, isLimit && d == up, true)
            }
        }
        return res % mod
    }
    return dfs(0, 0, true, false)
}

// 如果 low 是步进数字,那么多减了1,再加一不回来
func valid(s string) bool {
    for i := 1; i < len(s); i++ {
        if abs(int(s[i-1]), int(s[i])) != 1 {
            return false
        }
    }
    return true
}
