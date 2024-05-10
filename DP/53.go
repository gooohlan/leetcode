package DP

import "math"

func maxSubArrayDP(nums []int) int {
    dp, res := nums[0], nums[0]
    for i := 1; i < len(nums); i++ {
        dp = max(nums[i], dp+nums[i])
        if res < dp {
            res = dp
        }
    }
    return res
}

// 滑动窗口,窗口元素和大于0时扩大,小于0时重置窗口
// 因为和小于0就意味着窗口中存在了负数大于正数的情况,舍弃窗口中所有的数才为最佳
func maxSubArray(nums []int) int {
    sum, res := math.MinInt, math.MinInt
    for _, num := range nums {
        if sum > 0 {
            sum += num
        } else { // 舍弃窗口
            sum = num
        }
        
        if sum > res {
            res = sum
        }
    }
    return res
}

// 前缀和,求出前缀和数组
// 在前缀和数组中,以nums[i]结尾最大的子数组之和,就是preSum[i+1]- min(preSum[0...i])
func maxSubArrayPreSum(nums []int) int {
    preSum := make([]int, len(nums)+1)
    preSum[0] = nums[0]
    for i := 1; i <= len(nums); i++ {
        preSum[i] = preSum[i-1] + nums[i-1]
    }
    res := math.MinInt
    minVal := math.MaxInt // 记录0-i最小前缀和
    for i := 0; i < len(nums); i++ {
        minVal = min(minVal, preSum[i])
        res = max(res, preSum[i+1]-minVal)
    }
    return res
}
