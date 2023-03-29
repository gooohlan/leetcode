package DataStructure

import "sort"

func merge(intervals [][]int) [][]int {
    var res [][]int
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    res = append(res, intervals[0])
    for _, interval := range intervals[1:] {
        last := res[len(res)-1]
        if interval[0] > last[1] {
            res = append(res, interval)
        } else if interval[1] > last[1] {
            last[1] = interval[1]
        }
    }
    
    return res
}
