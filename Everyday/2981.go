package Everyday

import "sort"

func maximumLength1(s string) int {
    groups := make(map[byte][]int)
    cnt := 0
    for i, v := range s {
        cnt++
        if i+1 == len(s) || s[i+1] != byte(v) {
            groups[byte(v)] = append(groups[byte(v)], cnt)
            cnt = 0
        }
    }
    
    ans := 0
    for _, l := range groups {
        if len(l) == 0 {
            continue
        }
        sort.Slice(l, func(i, j int) bool {
            return l[i] > l[j]
        })
        
        l = append(l, 0, 0) // 添加两个空串,防止长度不够
        ans = max(ans, l[0]-2, min(l[0]-1, l[1]), l[2])
    }
    if ans == 0 {
        return -1
    }
    return ans
}
