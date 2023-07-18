package WeeklyContest

import "sort"

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
    type data struct {
        i, p, h int
        d       byte
    }
    
    arr := make([]data, len(positions))
    for i, position := range positions {
        arr[i] = data{i, position, healths[i], directions[i]}
    }
    sort.Slice(arr, func(i, j int) bool {
        return arr[i].p < arr[j].p
    })
    
    var toLeft, st []data
    for _, p := range arr {
        if p.d == 'R' {
            st = append(st, p)
            continue
        }
        
        for len(st) > 0 {
            top := &st[len(st)-1]
            if top.h > p.h { // 栈顶的健康度更大
                top.h--
                break
            }
            if top.h == p.h { // 健康度一样大
                st = st[:len(st)-1]
                p.h = 0 // 防止加入 toLeft
                break
            }
            p.h-- // p的健康度更大
            st = st[:len(st)-1]
        }
        if len(st) == 0 && p.h != 0 {
            toLeft = append(toLeft, p) // p把栈中全部撞掉
        }
    }
    
    toLeft = append(toLeft, st...)
    sort.Slice(toLeft, func(i, j int) bool {
        return toLeft[i].i < toLeft[j].i
    })
    res := make([]int, len(toLeft))
    for i, d := range toLeft {
        res[i] = d.h
    }
    return res
}
