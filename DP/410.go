package DP

import "math"

var memos [][]int

func splitArray(nums []int, k int) int {
    s := make([]int, len(nums)+1)
    for i := 1; i <= len(nums); i++ {
        s[i] = nums[i-1] + s[i-1]
    }
    
    memos = make([][]int, len(s))
    for i := 0; i < len(memos); i++ {
        memos[i] = make([]int, k+1)
        for j := 0; j < len(memos[i]); j++ {
            memos[i][j] = math.MaxInt32
        }
    }
    
    return dp410(0, k, s)
}

// [i, n)是待划分的区域，m表示划分为几段
func dp410(i, m int, s []int) int {
    if memos[i][m] != math.MaxInt32 {
        return memos[i][m]
    }
    if 1 == m { // 表示当前区域划分完毕
        return s[len(s)-1] - s[i]
    }
    
    minSum := math.MaxInt
    for k := i; k < len(s)-2; k++ {
        t := max(s[k+1]-s[i], dp410(k+1, m-1, s))
        minSum = min(t, minSum)
    }
    
    memos[i][m] = minSum
    return minSum
}

// dp[i][k] 表示数组前 i 个数分割为 k 段所得到的最大连续子数组和的最小值
// dp[j][k-1] 表示前 j 个数被分割为 k-1 段得到的最大连续子数组和的最小值
// 推导公式得出 dp[i][k] = max(dp[j][k−1],sum[j+1,i])
func splitArrayDP(nums []int, m int) int {
    n := len(nums)
    
    dp := make([][]int, n+1)
    sub := make([]int, n+1)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, m+1)
        for j := 0; j < len(dp[i]); j++ {
            dp[i][j] = math.MaxInt32
        }
    }
    
    for i := 1; i <= n; i++ {
        sub[i] = sub[i-1] + nums[i-1]
    }
    
    dp[0][0] = 0
    
    // 从分割为1个数组开始,到最多m结束
    for k := 1; k <= m; k++ {
        // i需要切割为k个数组,所以i最小值为k,最大值为数组长度
        for i := k; i <= n; i++ {
            // j要切割成k-1个数组, 并且j不能超过i
            for j := k - 1; j < i; j++ {
                dp[i][k] = min(dp[i][k], max(dp[j][k-1], sub[i]-sub[j]))
            }
        }
    }
    return dp[n][m]
}
