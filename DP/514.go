package DP

import "math"

func findRotateSteps(ring string, key string) int {
    indexMap := make([][]int, 26)
    for i, s := range ring {
        c := s - 'a'
        indexMap[c] = append(indexMap[c], i)
    }
    m, n := len(ring), len(key)
    var dfs func(i, j int) int // 圆盘指针在ring[i], key[j]最小操作数
    dfs = func(i, j int) int {
        if j == n {
            return 0
        }
        
        cur := key[j] - 'a'
        res := math.MaxInt
        for _, target := range indexMap[cur] {
            distance := abs(i - target)
            distance = min(distance, m-distance)        // 顺时针和逆时针选择一个最小的
            res = min(res, 1+distance+dfs(target, j+1)) // 调到ring[target],继续寻找key[k+1]
        }
        return res
    }
    
    return dfs(0, 0)
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
