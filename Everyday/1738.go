package Everyday

import (
    "sort"
)

func kthLargestValue(matrix [][]int, k int) int {
    m, n := len(matrix), len(matrix[0])
    sum := make([][]int, m+1)
    sum[0] = make([]int, n+1)
    arr := make([]int, 0)
    for i, row := range matrix {
        sum[i+1] = make([]int, n+1)
        for j, x := range row {
            sum[i+1][j+1] = sum[i+1][j] ^ sum[i][j+1] ^ sum[i][j] ^ x
        }
        arr = append(arr, sum[i+1][1:]...)
    }
    
    sort.Ints(arr)
    return arr[len(arr)-k]
}

func kthLargestValue2(matrix [][]int, k int) int {
    m, n := len(matrix), len(matrix[0])
    arr := make([]int, 0, m*n)
    colSum := make([]int, n)
    for _, row := range matrix {
        sum := 0
        for j, x := range row {
            colSum[j] ^= x
            sum ^= colSum[j]
            arr = append(arr, sum)
        }
    }
    
    sort.Ints(arr)
    return arr[len(arr)-k]
}
