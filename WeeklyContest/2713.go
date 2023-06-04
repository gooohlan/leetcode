package WeeklyContest

import "sort"

func maxIncreasingCells(mat [][]int) int {
    type pair struct {
        x, y int
    }
    g := make(map[int][]pair)
    for i, row := range mat {
        for j, x := range row {
            g[x] = append(g[x], pair{i, j})
        }
    }
    
    keys := make([]int, 0)
    for k := range g {
        keys = append(keys, k)
    }
    sort.Ints(keys)
    
    rowMax := make([]int, len(mat))
    colMax := make([]int, len(mat[0]))
    var ans int
    for _, x := range keys {
        pos := g[x]
        mx := make([]int, len(pos))
        for i, p := range pos {
            mx[i] = max(rowMax[p.x], colMax[p.y]) + 1
            ans = max(ans, mx[i])
        }
        for i, p := range pos {
            rowMax[p.x] = max(rowMax[p.x], mx[i])
            colMax[p.y] = max(colMax[p.y], mx[i])
        }
    }
    return ans
}
