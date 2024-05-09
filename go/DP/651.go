package DP

func maxA(n int) int {
    var dfs func(int, int, int) int
    
    dfs = func(n, nums, copy int) int {
        if n <= 0 {
            return nums
        }
        
        return max(
            dfs(n-1, nums+1, copy),    // A
            dfs(n-1, nums+copy, copy), // ctrl+V
            dfs(n-2, nums, nums),      // ctrl+c, ctrl+v
        )
    }
    
    return dfs(n, 0, 0)
}

func maxA2(n int) int {
    dp := make([]int, n+1)
    for i := 1; i <= n; i++ {
        dp[i] = dp[i-1] + 1
        for j := 2; j < i; j++ {
            // 全选 & 复制 dp[j-2]，连续粘贴 i - j 次
            // 屏幕上共 dp[j - 2] * (i - j + 1) 个 A
            dp[i] = max(dp[i], dp[j-2]*(i-j+1))
        }
    }
    
    return dp[n]
}
