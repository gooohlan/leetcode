package WeeklyContest

import "math"

func minCost(nums []int, x int) int64 {
    n := len(nums)
    sum := make([]int, n)
    for i := range sum {
        sum[i] = i * x // 操作i次的代价
    }
    
    // 这里循环为两个nums数组,将nums扩充一倍
    for i, num := range nums { // 子数组最左
        for j := i; j < n+i; j++ { // 子数组最右端
            num = min(num, nums[j%n]) // 从nums[i]到nums[j%n]的最小值
            sum[j-i] += num           // 位移j-i次取出num的成本
        }
    }
    
    ans := math.MaxInt
    for _, s := range sum {
        ans = min(s, ans)
    }
    return int64(ans)
}
