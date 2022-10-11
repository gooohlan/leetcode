package Stock

func maxProfit123(prices []int) int {
	var maxK = 2
	n := len(prices)
	var dp [][][2]int
	for i := 0; i < n; i++ {
		item := make([][2]int, maxK+1)
		dp = append(dp, item)
	}

	for i := 0; i < n; i++ {
		for k := maxK; k >= 1; k-- {
			if i == 0 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[n-1][maxK][0]
}

func maxProfit123_1(prices []int) int {
	dp10, dp11 := 0, -prices[0]
	dp20, dp21 := 0, -prices[0]
	for _, price := range prices {
		dp20 = max(dp20, dp21+price) // dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][2][1] + prices[i])
		dp21 = max(dp21, dp10-price) // dp[i][2][1] = max(dp[i-1][2][1], dp[i-1][1][0] - prices[i])
		dp10 = max(dp10, dp11+price) // dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1] + prices[i])
		dp11 = max(dp11, -price)     // dp[i][1][1] = max(dp[i-1][1][1], -prices[i])
	}
	return dp20
}
