package DP

func stoneGame(piles []int) bool {
    n := len(piles)
    // dp[i][j]：作为先手，在区间 nums[i..j] 里进行选择可以获得的相对分数
    dp := make([]int, n)
    // dp := make([][]int, n)
    for i := range dp {
        // dp[i] = make([]int, n)
        // dp[i][i] = nums[i]
        dp[i] = piles[i]
    }
    
    for i := n - 2; i >= 0; i-- {
        for j := i + 1; j < n; j++ {
            // 当前自己选择的分数为正,对手现在的为负数
            dp[j] = max(piles[i]-dp[j], piles[j]-dp[j-1])
        }
    }
    
    return dp[n-1] > 0
}
