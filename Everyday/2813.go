package Everyday

import "sort"

func findMaximumElegance(items [][]int, k int) int64 {
    sort.Slice(items, func(i, j int) bool {
        return items[i][0] > items[j][0]
    })
    
    ans, totalProfit := 0, 0
    vis := make(map[int]bool)
    var repeat []int
    
    for _, item := range items[:k] {
        profit, category := item[0], item[1]
        totalProfit += profit
        if !vis[category] {
            vis[category] = true
        } else {
            repeat = append(repeat, profit)
        }
    }
    
    for _, item := range items[k:] {
        profit, category := item[0], item[1]
        if len(repeat) > 0 && !vis[category] {
            vis[category] = true
            totalProfit = totalProfit - repeat[len(repeat)-1] + profit
            repeat = repeat[:len(repeat)-1]
        }
        
    }
    return int64(ans)
}
