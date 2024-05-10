package WeeklyContest

func countPartitions(nums []int, k int) int {
    const mod int = 1e9 + 7
    sum := 0
    for _, num := range nums {
        sum += num
    }
    if sum < k*2 {
        return 0
    }
    
    n := len(nums)
    dp := make([][]int, n+1) // 前i个数,和为j的方案数
    for i := range dp {
        dp[i] = make([]int, k)
    }
    dp[0][0] = 1
    res := 1
    for i := 1; i <= n; i++ {
        res = res * 2 % mod
        for j := 0; j < k; j++ {
            if j < nums[i-1] {
                dp[i][j] = dp[i-1][j]
            } else {
                dp[i][j] = (dp[i-1][j] + dp[i-1][j-nums[i-1]]) % mod
            }
        }
    }
    
    for _, v := range dp[n] {
        res = res - v*2
    }
    
    return (res%mod + mod) % mod
}
