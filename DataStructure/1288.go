package DataStructure

import "sort"

func removeCoveredIntervals(intervals [][]int) int {
    // 按照起点升序排列，起点相同时降序排列
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] > intervals[j][1]
        }
        return intervals[i][0] < intervals[j][0]
    })
    
    left, right := intervals[0][0], intervals[0][1]
    
    res := 0
    for _, interval := range intervals[1:] {
        // 起点升序,所以left <= interval[0]永远为true
        if left <= interval[0] && right >= interval[1] {
            res++
        }
        
        // 因为起点升序,终点降序,所以 right >= interval[0]永远为true
        if right >= interval[0] && right <= interval[1] {
            right = interval[1]
        }
        
        if right < interval[0] {
            left, right = interval[0], interval[1]
        }
    }
    
    return len(intervals) - res
}

func removeCoveredIntervals2(intervals [][]int) int {
    // 按照起点升序排列，起点相同时终点降序排列
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] == intervals[j][0] {
            return intervals[i][1] > intervals[j][1]
        }
        return intervals[i][0] < intervals[j][0]
    })
    
    right := intervals[0][1]
    
    res := 0
    for _, interval := range intervals[1:] {
        // 起点升序,所以left <= interval[0]永远为true
        if right >= interval[1] {
            res++
        }
        
        // 因为起点升序,终点降序,所以 right >= interval[0]永远为true
        if right <= interval[1] {
            right = interval[1]
        }
        
        if right < interval[0] {
            right = interval[1]
        }
    }
    
    return len(intervals) - res
}
