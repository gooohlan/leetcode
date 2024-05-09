package DP

import "sort"

// 把原数组根据邮箱的W进行递增排序,然后对于h像题300那样处理
func maxEnvelopes(envelopes [][]int) int {
    // 按宽度升序排列，如果宽度一样，则按高度降序排列
    sort.Slice(envelopes, func(i, j int) bool {
        if envelopes[i][0] == envelopes[j][0] {
            return envelopes[i][1] > envelopes[j][1]
        }
        return envelopes[i][0] < envelopes[j][0]
    })
    
    arr := make([]int, len(envelopes))
    for i, envelope := range envelopes {
        arr[i] = envelope[1]
    }
    
    return lengthOfLIS2(arr)
}
