package Stock

// 本体与122最大的区别就是买入卖出多次,所以在买入时做一些调整
func maxProfit122(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < n; i++ {
		// dp[i][0] 表示在第I天未持有, 此时有两种情况:昨天就未持有, 或者昨天持有今日卖了,二者取最大值即可
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// dp[i][1] 表示在第I天持有, 此时有两种情况:昨天就持有, 或者昨天未持有今日买进,买进时需要从最大利润中减去今天买入价格,二者取最大值即可
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[n-1][0]
}

// 一样的,依旧可以压缩空间
func maxProfit122_1(prices []int) int {
	dp0, dp1 := 0, -prices[0]

	for i := 1; i < len(prices); i++ {
		dp0, dp1 = max(dp0, dp1+prices[i]), max(dp1, dp0-prices[i])
	}

	return dp0
}

// 这里还提供一个更简单的思路,因为是可以重复买入买出,所以当第今天的价格小于明天的就在今天买入明天卖出,依次类推
func maxProfit122_2(prices []int) int {
	max := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			max += prices[i+1] - prices[i]
		}
	}
	return max
}
