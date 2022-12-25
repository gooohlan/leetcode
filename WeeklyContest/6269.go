package WeeklyContest

import "math"

func closetTarget(words []string, target string, startIndex int) int {
    n := len(words)
    res := math.MaxInt
    for i, word := range words {
        if word == target {
            res = min(res, abs(startIndex, i))   // 直接移动
            res = min(res, n-abs(startIndex, i)) // 绕环移动
        }
    }
    if res == math.MaxInt {
        return -1
    }
    return res
}
