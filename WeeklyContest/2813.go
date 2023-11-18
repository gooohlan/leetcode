package WeeklyContest

import "sort"

// 利润从大到小排序, 先选择前 k 个项目,计算最大利润
// 遍历 [k+1,n)中的项目, 只有当第 k+i 个项目为选择过,且前面选择的项目没有重复时才能从原来选择的项目中替换以增加利润,因为:
// - 如果 k+i 个项目与前面的类别相同,那么他的利润更小,不需要选他
// - 如果前面项目为出现重复, 那么类别一增一减 distinct_categories 不会有变化,而后面的利润比前面小, 反而会降低
// 我们需要保证 total_profit 劲量大的同时去增加 categories 的数量,以达到目的.
func findMaximumElegance(items [][]int, k int) int64 {
    sort.Slice(items, func(i, j int) bool {
        return items[i][0] > items[j][0]
    })
    ans, totalProfit := 0, 0
    vis := map[int]bool{}
    repeat := []int{}
    
    for _, item := range items[:k] {
        profit, category := item[0], item[1]
        totalProfit += profit
        if !vis[category] {
            vis[category] = true
        } else {
            repeat = append(repeat, profit)
        }
    }
    
    ans = totalProfit + len(vis)*len(vis)
    for _, item := range items[k:] {
        profit, category := item[0], item[1]
        if len(repeat) > 0 && !vis[category] { // 类别为出现过,且前面出现过重复
            vis[category] = true
            totalProfit += profit - repeat[len(repeat)-1] // 替换利润最小的
            repeat = repeat[:len(repeat)-1]
        }
        ans = max(ans, totalProfit+len(vis)*len(vis))
    }
    return int64(ans)
}
