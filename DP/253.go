package DP

import "sort"

func minMeetingRooms(intervals [][]int) int {
    n := len(intervals)
    begin, end := make([]int, n), make([]int, n)
    for i, interval := range intervals {
        begin[i] = interval[0]
        end[i] = interval[1]
    }
    
    sort.Slice(begin, func(i, j int) bool {
        return begin[i] < begin[j]
    })
    sort.Slice(end, func(i, j int) bool {
        return end[i] < end[j]
    })
    
    count, res := 0, 0
    i, j := 0, 0 // i扫描头结点,j扫描尾节点
    
    for i < n && j < n {
        if begin[i] < end[j] {
            // 扫描到头结点
            count++
            i++
        } else {
            count--
            j++
        }
        res = max(res, count)
    }
    return res
}
