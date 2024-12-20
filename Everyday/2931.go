package Everyday

import "sort"

func maxSpending(values [][]int) int64 {
    m, n := len(values), len(values[0])
    arr := make([]int, 0, m*n)
    for _, value := range values {
        arr = append(arr, value...)
    }
    sort.Ints(arr)
    
    var res int64
    for i, x := range arr {
        res += int64(x) * int64(i+1)
    }
    return res
}
