package Everyday

func findTargetSumWays(nums []int, target int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    sum -= abs(0, target)
    // p−q=target p+q=s, 2q=s-target
    if sum < 0 || sum%2 == 1 {
        return 0
    }
    
    m := sum / 2 // 背包数量
    n := len(nums)
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, m+1)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    
    var dfs func(int, int) int
    dfs = func(i int, c int) int {
        if i < 0 {
            if c == 0 {
                return 1
            }
            return 0
        }
        
        if memo[i][c] == -1 {
            if c < nums[i] {
                memo[i][c] = dfs(i-1, c)
            } else {
                memo[i][c] = dfs(i-1, c) + dfs(i-1, c-nums[i])
            }
        }
        return memo[i][c]
    }
    return dfs(n-1, m)
}

func findTargetSumWays2(nums []int, target int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    sum -= abs(0, target)
    // p−q=target p+q=s, 2q=s-target
    if sum < 0 || sum%2 == 1 {
        return 0
    }
    
    m := sum / 2 // 背包数量
    n := len(nums)
    
    dp := make([][]int, 2)
    for i := range dp {
        dp[i] = make([]int, m+1)
    }
    dp[0][0] = 1
    
    for i, num := range nums {
        for c := 0; c <= m; c++ {
            if c < num {
                // dp[i+1][c] = dp[i][c]
                dp[(i+1)%2][c] = dp[i%2][c]
            } else {
                // dp[i+1][c] = dp[i][c] + dp[i][c-num]
                dp[(i+1)%2][c] = dp[i%2][c] + dp[i%2][c-num]
            }
        }
    }
    return dp[n][m]
    // return dp[n%2][m]
}
func findTargetSumWays3(nums []int, target int) int {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    sum -= abs(0, target)
    // p−q=target p+q=s, 2q=s-target
    if sum < 0 || sum%2 == 1 {
        return 0
    }
    
    m := sum / 2 // 背包数量
    
    dp := make([]int, m+1)
    dp[0] = 1
    
    for _, num := range nums {
        for c := m; c >= num; c-- {
            dp[c] += dp[c-num]
        }
    }
    return dp[m]
}
