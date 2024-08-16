package Everyday

import "math"

func maxScore3184(grid [][]int) int {
    ans := math.MinInt
    m, n := len(grid), len(grid[0])
    f := make([][]int, m+1)
    f[0] = make([]int, n+1)
    for i := range f[0] {
        f[0][i] = math.MaxInt
    }
    
    for i, row := range grid {
        f[i+1] = make([]int, n+1)
        f[i+1][0] = math.MaxInt
        for j, x := range row {
            minX := min(f[i+1][j], f[i][j+1])
            ans = max(ans, x-minX)
            f[i+1][j+1] = min(minX, x)
        }
    }
    return ans
}

func maxScore3184_2(grid [][]int) int {
    ans := math.MinInt
    colMin := make([]int, len(grid[0]))
    for i := range colMin {
        colMin[i] = math.MaxInt
    }
    
    for i, row := range grid {
        preMin := math.MaxInt // 这一行最小值,也是 colMin[0,j]最小值
        for j, x := range row {
            ans = max(ans, x-min(preMin, colMin[j]))
            colMin[j] = min(colMin[j], x)   // 维护每一列最小值
            preMin = min(preMin, colMin[i]) // 维护当前行最小值
        }
    }
    return ans
}
