package DP

import "sort"

func videoStitching(clips [][]int, time int) (ans int) {
    // 按照起点升序排序,起点相同就用降序排列
    sort.Slice(clips, func(i, j int) bool {
        if clips[i][0] == clips[j][0] {
            return clips[i][1] > clips[j][1]
        }
        return clips[i][0] < clips[j][0]
    })
    
    res := 0
    curEnd, nextEnd := 0, 0
    i, n := 0, len(clips)
    for i < n && clips[i][0] <= curEnd {
        // 在第 res 个视频的区间内贪心选择下一个视频
        for i < n && clips[i][0] <= curEnd {
            nextEnd = max(nextEnd, clips[i][1])
            i++
        }
        // 找到下一个视频,更新curEnd
        res++
        curEnd = nextEnd
        if curEnd >= time {
            return res
        }
    }
    return -1
}
