package WeeklyContest

import (
    "math"
    "sort"
)

func maxPower(stations []int, r int, k int) int64 {
    n := len(stations)
    // 前缀和
    sums := make([]int, n+1)
    for i, x := range stations {
        sums[i+1] = sums[i] + x
    }
    
    // 求出每个城市初始电量
    mn := math.MaxInt // 最小耗电量城市的电量
    for i := range stations {
        stations[i] = sums[min(i+r+1, n)] - sums[max(i-r, 0)] // i处的总电量为[i-r, i+r]
        mn = min(mn, stations[i])
    }
    
    // 在哪儿建电站
    // 建在min(i+r, n-1)
    // 影响范围[i, min(i+2r,n)]
    // var check func(x int) bool
    // check = func(minPower int) bool {
    //     diff := make([]int, n) // 差分数组,用于安装电站
    //     sumD, need := 0, 0     // sumD累加求差分和,更新差分;
    //     for i, power := range stations {
    //         sumD += diff[i]
    //         m := minPower - power - sumD // 需要建造的供电站
    //         if m > 0 {
    //             need += m
    //             if need > k { // 需要建造的供电站大于k,直接返回失败
    //                 return false
    //             }
    //             sumD += m // 更新差分和,建造供电站
    //             if i+r*2+1 < n {
    //                 diff[i+r*2+1] -= m // 差分更新
    //             }
    //         }
    //     }
    //     return true
    // }
    //
    // left, right := mn, sums[n]+k+1
    //
    // for left+1 < right {
    //     mid := (left + right) / 2
    //     if check(mid) {
    //         left = mid
    //     } else {
    //         right = mid
    //     }
    // }
    // return int64(left)
    return int64(mn + sort.Search(k, func(minPower int) bool {
        minPower += mn + 1 // 改为二分最小的不满足要求的值，这样 sort.Search 返回的就是最大的满足要求的值
        sumD, need := 0, 0
        diff := make([]int, n) // 差分数组
        for i, power := range stations {
            sumD += diff[i]
            m := minPower - power - sumD
            if m > 0 {
                need += m
                if need > k {
                    return true // 与手写相反
                }
                sumD += m
                if i+2*r+1 < n {
                    diff[i+2*r+1] -= m
                }
            }
        }
        return false
    }))
}
