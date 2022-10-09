package Stock

// 与前两题思路大体一致,区别在第i天买入时,要在i-2天买出的基础上操作,因为有一天冻结期
func maxProfit309(prices []int) int {
	n := len(prices)
	var dp [][2]int

	for i := 0; i < n; i++ {
		var item [2]int
		if i == 0 {
			item[0] = 0
			item[1] = -prices[i]
		} else if i == 1 {
			item[0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
			item[1] = max(dp[i-1][1], -prices[i]) // 避免i-2小于0
		} else {
			item[0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
			item[1] = max(dp[i-1][1], dp[i-2][0]-prices[i]) // 第i-1天冷冻
		}
		dp = append(dp, item)
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
