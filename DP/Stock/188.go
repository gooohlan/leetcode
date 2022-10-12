package Stock

// 处理123时我们定义一个maxK,这里最大交易次数直接复用即可
// 这里有小细节,k不能是无限打的,他应该有所约束,如果当k>n/2(买卖不能同时进行,所以n/2)时, k就没有约束效果,相当于没有最多k次交易的约束

func maxProfit188(k int, prices []int) int {
	n := len(prices)
	maxK := k

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
	return dp[n-1][k][0]
}
