package DP

import "sort"

func findMinArrowShots(points [][]int) int {
    n := len(points)
    if n == 0 {
        return 0
    }
    
    sort.Slice(points, func(i, j int) bool {
        return points[i][1] < points[j][1]
    })
    
    count := 1
    right := points[0][1]
    
    for _, point := range points[1:] {
        if point[0] > right {
            count++
            right = point[1]
        }
    }
    
    return count
}
