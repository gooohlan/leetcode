package WeeklyContest

import (
    "container/heap"
    "sort"
)

// 思路参考 https://leetcode.cn/problems/time-to-cross-a-bridge/solution/by-endlesscheng-nzqo/
func findCrossingTime(n int, k int, time [][]int) int {
    // 首先排序,效率从高到低排序
    sort.SliceStable(time, func(i, j int) bool {
        a, b := time[i], time[j]
        return a[0]+a[2] < b[0]+b[2]
    })
    
    waitL, waitR := make(hp6284, k), hp6284{} // waitR：右边等待过桥的工人 waitL：左边等待过桥的工人
    for i := range waitL {
        waitL[i].i = k - 1 - i // 下标越大效率越低
    }
    
    workL, workR := hp62842{}, hp62842{} // workL：新仓库正在放箱的工人; workR：旧仓库正在搬箱的工人
    
    cur := 0
    for n > 0 {
        for len(workL) > 0 && workL[0].t <= cur { // 把左边完成放箱的工人放入左边等待队列
            heap.Push(&waitL, heap.Pop(&workL))
        }
        
        for len(workR) > 0 && workR[0].t <= cur { // 右边完成搬箱工人放入右边等待队列
            heap.Push(&waitR, heap.Pop(&workR))
        }
        
        if len(waitR) > 0 { // 右边过桥，注意加到 waitR 中的都是 <= cur 的（下同）
            p := heap.Pop(&waitR).(pair6284)
            cur += time[p.i][2]
            heap.Push(&workL, pair6284{p.i, cur + time[p.i][3]}) // 放箱，记录完成时间
        } else if len(waitL) > 0 { // 左边过桥
            p := heap.Pop(&waitL).(pair6284)
            cur += time[p.i][0]
            heap.Push(&workR, pair6284{p.i, cur + time[p.i][1]}) // 搬箱，记录完成时间
            n--
        } else if len(workL) == 0 { // cur 过小，找个最小的放箱/搬箱完成时间来更新 cur
            cur = workR[0].t
        } else if len(workR) == 0 {
            cur = workL[0].t
        } else {
            cur = min(workR[0].t, workL[0].t)
        }
    }
    
    for len(workR) > 0 {
        p := heap.Pop(&workR).(pair6284) // 右边完成搬箱
        // 如果没有排队，直接过桥；否则由于无论谁先过桥，最终完成时间都一样，所以也可以直接计算
        cur = max(p.t, cur) + time[p.i][2]
    }
    
    return cur
}

type pair6284 struct{ i, t int }
type hp6284 []pair6284

func (h hp6284) Len() int            { return len(h) }
func (h hp6284) Less(i, j int) bool  { return h[i].i > h[j].i }
func (h hp6284) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp6284) Push(v interface{}) { *h = append(*h, v.(pair6284)) }
func (h *hp6284) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

type hp62842 []pair6284

func (h hp62842) Len() int            { return len(h) }
func (h hp62842) Less(i, j int) bool  { return h[i].t < h[j].t }
func (h hp62842) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp62842) Push(v interface{}) { *h = append(*h, v.(pair6284)) }
func (h *hp62842) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
