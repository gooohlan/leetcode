package Everyday

import (
    "sort"
)

func maxIncreasingCells(mat [][]int) int {
    g := make(map[int][][2]int)
    for i, row := range mat {
        for j, x := range row {
            g[x] = append(g[x], [2]int{i, j})
        }
    }
    
    keys := make([]int, 0, len(g))
    for k := range g {
        keys = append(keys, k)
    }
    sort.Ints(keys)
    
    rowMax := make([]int, len(mat))
    colMax := make([]int, len(mat[0]))
    
    var ans int
    for _, k := range keys {
        pos := g[k]
        fs := make([]int, len(pos))
        for i, p := range pos {
            fs[i] = max(rowMax[p[0]], colMax[p[1]]) + 1
            ans = max(ans, fs[i])
        }
        
        for i, p := range pos {
            rowMax[p[0]] = max(rowMax[p[0]], fs[i])
            colMax[p[1]] = max(colMax[p[1]], fs[i])
        }
    }
    return ans
}
