package Stock

// 交易存在手续费,那就从卖出的中减去即可
func maxProfit714(prices []int, fee int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

func maxProfit714_1(prices []int, fee int) int {
	dp0, dp1 := 0, -prices[0]
	for i := 0; i < len(prices); i++ {
		dp0, dp1 = max(dp0, dp1+prices[i]-fee), max(dp1, dp0-prices[i])
	}
	return dp0
}
