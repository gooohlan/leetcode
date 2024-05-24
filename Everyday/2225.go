package Everyday

import "sort"

func findWinners(matches [][]int) [][]int {
    m := make(map[int]int)
    for _, match := range matches {
        if m[match[0]] == 0 {
            m[match[0]] = 0
        }
        
        m[match[1]]++
    }
    
    res := make([][]int, 2)
    for user, count := range m {
        if count < 2 {
            res[count] = append(res[count], user)
        }
    }
    
    sort.Ints(res[0])
    sort.Ints(res[1])
    return res
}
