package Stock

// 与前两题思路大体一致,区别在第i天买入时,要在i-2天买出的基础上操作,因为有一天冻结期
func maxProfit309(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	if n > 1 {
		dp[1][0] = max(dp[0][0], dp[0][1]+prices[1])
		dp[1][1] = max(dp[0][1], -prices[1])

		for i := 2; i < n; i++ {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
			dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i]) // 第i-1天冷冻
		}
	}

	return dp[n-1][0]
}

func maxProfit309_1(prices []int) int {
	dp0, dp1, dp_pre := 0, -prices[0], -prices[0]
	// dp_pre 代表 dp[i-2][0]
	for i := 0; i < len(prices); i++ {
		dp0, dp1, dp_pre = max(dp0, dp1+prices[i]), max(dp1, dp_pre-prices[i]), dp0
	}

	return dp0
}
