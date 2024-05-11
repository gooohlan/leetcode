package Everyday

import "strings"

// 先假设每辆车都走到最后,然后倒序循环每种垃圾,如果没有这种垃圾,减去前面加上的时间
func garbageCollection(garbage []string, travel []int) int {
    var res int
    // 收集垃圾总用时
    for _, s := range garbage {
        res += len(s)
    }

    // 假设每辆车都走到了最后
    for _, t := range travel {
        res += t * 3
    }

    for _, c := range []byte("MPG") {
        for i := len(garbage) - 1; i > 0 && strings.IndexByte(garbage[i], c) == -1; i-- {
            res -= travel[i-1] // 没有垃圾C,减去前面加上的时间
        }
    }
    return res
}

// 使用map记录每辆车行驶用时,再记录一个总耗时,当前房间存在该垃圾,记录用时车辆就添加
func garbageCollection1(garbage []string, travel []int) int {
    m := make(map[rune]int)
    m['M'] = 0
    m['P'] = 0
    m['G'] = 0

    res := len(garbage[0])
    sumTime := 0

    for i, s := range garbage[1:] {
        res += len(s)        // 收集当前房间耗时
        sumTime += travel[i] // 移动到当前房间总耗时
        for _, c := range s {
            res += sumTime - m[c] // 收集垃圾c车移动耗时
            m[c] = sumTime
        }
    }
    return res
}

// 贡献法
// 从后往前遍历,假设最后一个房间出现了M,那么我们加上他走的最后一个房间的时间 travel[n-2]的时间
// 假设倒数第二个房间存在PG,此时先加上M走travel[n-3]的时间,再加上P和G走travel[n-3]的时间,依此类推
func garbageCollection2(garbage []string, travel []int) int {
    res := len(garbage[0])
    m := map[rune]bool{}
    for i := len(garbage) - 1; i > 0; i-- {
        g := garbage[i]
        for _, c := range g {
            m[c] = true
        }
        res += len(g) + travel[i-1]*len(m)
    }
    return res
}
