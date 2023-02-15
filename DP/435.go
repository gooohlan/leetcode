package DP

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
    n := len(intervals)
    if n == 0 {
        return 0
    }
    
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] < intervals[j][1]
    })
    count := 1 // 至少有一个区间不想交
    right := intervals[0][1]
    
    for _, item := range intervals[1:] {
        if item[0] >= right { // 区间左侧在上个区间右侧,不相交
            count++
            right = item[1]
        }
    }
    return n - count // 减去不相交区间
}

func eraseOverlapIntervalsDP(intervals [][]int) int {
    n := len(intervals)
    if n == 0 {
        return 0
    }
    
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    dp := make([]int, n)
    for i := range dp {
        dp[i] = 1
    }
    for i := 1; i < n; i++ {
        for j := 0; j < i; j++ {
            if intervals[j][1] <= intervals[i][0] { // 区间j与区间i不重合
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
    }
    return n - max(dp...)
}
