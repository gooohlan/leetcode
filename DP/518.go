package DP

func change(amount int, coins []int) int {
    n := len(coins)
    dp := make([][]int, n+1) // 表示使用前i个金币凑出金额为j的方法
    for i := range dp {
        dp[i] = make([]int, amount+1)
        dp[i][0] = 1
    }
    
    for i := 1; i <= n; i++ {
        for j := 1; j <= amount; j++ {
            if j >= coins[i-1] {
                // 不放入面值为coins[i-1]的金币,就继承dp[i-1][j], 放入则为dp[i][j-coins[i-1]], 因为求得是共有多少种,所以把两个相加
                dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
            } else {
                dp[i][j] = dp[i-1][j]
            }
        }
    }
    return dp[n][amount]
}

// 二维数组优化为一维数组
// 因为dp[i][j]只与dp[i-1][j]有关,所以可以优化
func change2(amount int, coins []int) int {
    dp := make([]int, amount+1) // 抽出金额为i的金币的次数
    dp[0] = 1
    
    for i := 0; i < len(coins); i++ {
        for j := coins[i]; j <= amount; j++ { // 完全背包需要从前往后遍历
            dp[j] += dp[j-coins[i]]
        }
    }
    
    return dp[amount]
}
