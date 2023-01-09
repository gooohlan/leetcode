package WeeklyContest

import (
    "container/heap"
    "sort"
)

func maxKelements(nums []int, k int) int64 {
    pq := hp6285{nums}
    heap.Init(&pq)
    var res int64
    for ; k > 0; k-- {
        res += int64(pq.IntSlice[0])
        pq.IntSlice[0] = (pq.IntSlice[0] + 2) / 3
        heap.Fix(&pq, 0)
    }
    
    return res
}

type hp6285 struct{ sort.IntSlice }

func (h hp6285) Less(i, j int) bool {
    return h.IntSlice[i] > h.IntSlice[j]
}
func (h hp6285) Push(interface{}) {}
func (h hp6285) Pop() (_ interface{}) {
    return
}
