package Everyday

func maxScore(nums []int, x int) int64 {
    n := len(nums)
    memo := make([][2]int, n)
    for i := range memo {
        memo[i] = [2]int{-1, -1}
    }
    
    var dfs func(int, int) int
    dfs = func(i int, j int) (res int) {
        if i == n {
            return
        }
        if memo[i][j] != -1 {
            return memo[i][j]
        }
        
        if nums[i]%2 != j {
            res = dfs(i+1, j)
        } else {
            res = max(dfs(i+1, j), dfs(i+1, j^1)-x) + nums[i]
        }
        memo[i][j] = res
        return res
    }
    
    return int64(dfs(0, nums[0]%2))
}

func maxScore2(nums []int, x int) int64 {
    n := len(nums)
    dp := make([][2]int, n+1)
    for i := n - 1; i >= 0; i-- {
        v := nums[i]
        r := v % 2
        dp[i][r^1] = dp[i+1][r^1]
        dp[i][r] = max(dp[i+1][r], dp[i+1][r^1]-x) + v
    }
    return int64(dp[0][nums[0]%2])
}
