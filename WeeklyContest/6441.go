package WeeklyContest

import "strconv"

func punishmentNumber(n int) int {
    var sum int
    
    for i := 1; i <= n; i++ {
        s := strconv.Itoa(i * i)
        n := len(s)
        var dfs func(int, int) bool
        dfs = func(p int, sum int) bool {
            if p == n { // 切割到最后一位,判断是否符合要求
                return sum == i
            }
            x := 0
            for j := p; j < n; j++ {
                x = x*10 + int(s[j]-'0')
                if dfs(j+1, sum+x) {
                    return true
                }
            }
            return false
        }
        if dfs(0, 0) {
            sum += i * i
        }
    }
    return sum
}
