package WeeklyContest

import "math"

func maxSum(nums []int) int {
    res := -1
    maxVal := [10]int{}
    for i := range maxVal {
        maxVal[i] = math.MinInt
    }
    
    for _, v := range nums {
        maxD := 0
        for x := v; x > 0; x /= 10 {
            maxD = max(maxD, x%10)
        }
        res = max(res, v+maxVal[maxD])
        maxVal[maxD] = max(maxVal[maxD], v)
    }
    return res
}
