package WeeklyContest

import (
    "math"
    
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
