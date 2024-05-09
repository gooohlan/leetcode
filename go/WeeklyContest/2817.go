package WeeklyContest

import (
    "math"
    "sort"
    
    "github.com/emirpasic/gods/trees/redblacktree"
)

func minAbsoluteDifference(nums []int, x int) int {
    res := math.MinInt
    t := redblacktree.NewWithIntComparator()
    t.Put(math.MaxInt, nil)
    t.Put(math.MinInt/2, nil)
    
    for i, y := range nums[x:] {
        t.Put(nums[i], nil)
        c, _ := t.Ceiling(y)
        f, _ := t.Floor(y)
        res = min(res, min(abs(c.Key.(int), y), abs(f.Key.(int), y)))
    }
    return res
}

// 从第 x 个数开始往后循环
// 每次往前面的有序数组中插入,然后寻找最近的数求最小绝对值
func minAbsoluteDifference2(nums []int, x int) int {
    res := math.MaxInt
    ordered := []int{}
    for i, y := range nums[x:] {
        ordered = InsertOrder(ordered, nums[i])
        idx := sort.SearchInts(ordered, y)
        if idx < len(ordered) {
            res = min(res, abs(ordered[idx], y))
        }
        if idx > 0 || idx == len(ordered) {
            res = min(res, abs(ordered[idx-1], y))
        }
    }
    return res
}

func InsertOrder(ordered []int, num int) []int {
    idx := sort.SearchInts(ordered, num)
    ordered = append(ordered[:idx], append([]int{num}, ordered[idx:]...)...)
    return ordered
}
