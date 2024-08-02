package Everyday

import (
    "slices"
)

func maxmiumScore(cards []int, cnt int) int {
    slices.SortFunc(cards, func(a, b int) int { return b - a })
    s := 0
    for i := 0; i < cnt; i++ {
        s += cards[i]
    }
    if s%2 == 0 {
        return s
    }
    
    replaceNum := func(x int) int {
        for _, v := range cards[cnt:] {
            if v%2 != x%2 { // 找到一个与x奇偶性不同的数
                return s - x + v
            }
        }
        return 0
    }
    x := cards[cnt-1]               // 现在选择的最小的数
    ans := replaceNum(x)            // 替换最小的数
    for i := cnt - 2; i >= 0; i-- { // 比x小的数
        if cards[i]%2 != x%2 { // 找到一个最小的奇偶性和 x 不同的数
            ans = max(ans, replaceNum(cards[i])) // 替换
            break
        }
    }
    return ans
}
