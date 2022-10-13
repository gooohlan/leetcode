package DP

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 从前往后抢,dp[i]表示到第i个房间最多可以抢到去的钱
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums)+2)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

func rob1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp1, dp2 := nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp1, dp2 = dp2, max(dp1+nums[i], dp2)
	}
	return dp2
}

// 从后往前抢，dp[i]表示从第 i 间房子开始抢劫，最多能抢到的钱
func rob2(nums []int) int {
	dp := make([]int, len(nums)+2)

	for i := len(nums) - 1; i >= 0; i-- {
		dp[i] = max(dp[i+1], nums[i]+dp[i+2])
	}
	return dp[0]
}

func rob3(nums []int) int {
	dp1, dp2 := 0, 0
	for i := len(nums) - 1; i >= 0; i-- {
		dp1, dp2 = max(dp1, nums[i]+dp2), dp1
	}
	return dp1
}
