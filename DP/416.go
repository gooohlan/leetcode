package DP

func canPartition(nums []int) bool {
    n := len(nums)
    var sum int
    for _, num := range nums {
        sum += num
    }
    if sum%2 != 0 {
        return false
    }
    sum /= 2
    
    dp := make([][]bool, n+1)
    for i := range dp {
        dp[i] = make([]bool, sum+1)
        dp[i][0] = true
    }
    
    for i := 1; i <= n; i++ {
        for j := 1; j <= sum; j++ {
            if j-nums[i-1] < 0 {
                dp[i][j] = dp[i-1][j]
            } else {
                dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
            }
        }
    }
    
    return dp[n][sum]
}
