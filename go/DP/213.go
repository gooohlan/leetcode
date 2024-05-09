package DP

// 在198题的基础上,当房间数大于2时会出现两种情况:
// 抢了第一间房,那么最后一间房不能抢
// 抢了最后一间房, 第一间房不能抢
func rob213(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(robRange(nums[:len(nums)-1]), robRange(nums[1:]))
}

func robRange(nums []int) int {
	dp1, dp2 := 0, 0
	for i := len(nums) - 1; i >= 0; i-- {
		dp1, dp2 = max(dp1, nums[i]+dp2), dp1
	}
	return dp1
}
