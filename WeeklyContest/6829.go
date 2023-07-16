package WeeklyContest

import (
    "sort"
)

// 1,2,4,6,4
func maximumBeauty(nums []int, k int) int {
    n := len(nums)
    
    sort.Ints(nums)
    d := make([]int, nums[n-1]+k+2)
    for _, num := range nums {
        d[max(0, num-k)] += 1
        d[num+k+1] -= 1
    }
    
    res := 0
    for i := max(0, nums[0]-k); i < len(d); i++ {
        res = max(res, res+d[i])
    }
    
    return res
}
