package Stock

// 此题的变化只有两个,一个是天数,第二个是当前的持有状态
// 声明一个二维数组 分别表示天数和是否购买, 1表示持有 0未持有
func maxprofit121(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	for i := 1; i < n; i++ {
		// dp[i][0] 表示在第I天未持有, 此时有两种情况:昨天就未持有, 或者昨天持有今日卖了,二者取最大值即可
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// dp[i][0] 表示在第I天持有, 此时有两种情况:昨天就持有, 或者昨天未持有今日买进,二者取最大值即可
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}

	return dp[n-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 从上面程序可以看出,每次状态的变化只与上一次的状态有关,可以压缩一下空间
func maxProfit121_1(prices []int) int {
	dp0, dp1 := 0, -prices[0]
	for i := 0; i < len(prices); i++ {
		dp0 = max(dp0, dp1+prices[i])
		dp1 = max(dp1, -prices[i])
	}
	return dp0
}
