package main

func findLengthOfLCIS(nums []int) int {
	var ans, start int
	for k, v := range nums {
		if k > 0 && v <= nums[k-1] {
			start = k
		}
		ans = max(ans, k-start +1)
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
