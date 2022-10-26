package DP

import "math"

func splitArray(nums []int, k int) int {
	s := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		s[i] = nums[i-1] + s[i-1]
	}
	return dp410(0, k, s)
}

// [i, n)是待划分的区域，j表示当前是第几个划分点，j个划分点能够划出j+1个区域
func dp410(i, m int, s []int) int {
	if 1 == m { // 表示当前区域划分完毕
		return s[len(s)-1] - s[i]
	}

	minSum := math.MaxInt
	for k := i; k < len(s)-2; k++ {
		t := max(s[k+1]-s[i], dp410(k+1, m-1, s))
		minSum = min(t, minSum)
	}
	return minSum
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
