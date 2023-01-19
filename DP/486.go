package DP

func PredictTheWinner(nums []int) bool {
    var dfs func(int, int, int) int
    
    dfs = func(first, second, turn int) int {
        if first == second {
            return nums[first] * turn
        }
        
        scoreFirst := nums[first]*turn + dfs(first+1, second, -turn)
        scoreSecond := nums[second]*turn + dfs(first, second-1, -turn)
        return max(scoreFirst*turn, scoreSecond*turn) * turn
    }
    
    return dfs(0, len(nums)-1, 1) >= 0
}

func PredictTheWinnerDP(nums []int) bool {
    n := len(nums)
    // dp[i][j]：作为先手，在区间 nums[i..j] 里进行选择可以获得的相对分数
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
        dp[i][i] = nums[i]
    }
    
    for i := n - 2; i >= 0; i-- {
        for j := i + 1; j < n; j++ {
            // 当前自己选择的分数为正,对手现在的为负数
            dp[i][j] = max(nums[i]-dp[i+1][j], nums[j]-dp[i][j-1])
        }
    }
    
    return dp[0][n-1] >= 0
}

func PredictTheWinnerDP2(nums []int) bool {
    n := len(nums)
    // dp[i][j]：作为先手，在区间 nums[i..j] 里进行选择可以获得的相对分数
    dp := make([]int, n)
    // dp := make([][]int, n)
    for i := range dp {
        // dp[i] = make([]int, n)
        // dp[i][i] = nums[i]
        dp[i] = nums[i]
    }
    
    for i := n - 2; i >= 0; i-- {
        for j := i + 1; j < n; j++ {
            // 当前自己选择的分数为正,对手现在的为负数
            dp[j] = max(nums[i]-dp[j], nums[j]-dp[j-1])
        }
    }
    
    return dp[n-1] >= 0
}
