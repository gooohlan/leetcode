package WeeklyContest

func count(num1 string, num2 string, minSum int, maxSum int) int {
    const mod int = 1e9 + 7
    f := func(s string) int {
        // 剪枝
        memo := make([][]int, len(s))
        for i := range memo {
            memo[i] = make([]int, min(9*len(s), maxSum)+1)
            for j := range memo[i] {
                memo[i][j] = -1
            }
        }
        
        var dfs func(i int, sum int, limitUp bool) int
        dfs = func(i int, sum int, limitUp bool) int {
            if sum > maxSum {
                return 0
            }
            if i == len(s) {
                if sum >= minSum {
                    return 1
                }
                return 0
            }
            var res int
            if !limitUp && memo[i][sum] != -1 {
                return memo[i][sum]
            }
            
            up := 9
            if limitUp {
                up = int(s[i] - '0')
            }
            for d := 0; d <= up; d++ {
                res = (res + dfs(i+1, sum+d, limitUp && d == up)) % mod
            }
            if !limitUp {
                memo[i][sum] = res
            }
            return res
        }
        return dfs(0, 0, true)
    }
    ans := f(num2) - f(num1) + mod // 避免负数
    sum := 0
    for _, c := range num1 {
        sum += int(c - '0')
    }
    if minSum <= sum && sum <= maxSum { // x=num1 是合法的，补回来
        ans++
    }
    return ans % mod
}
