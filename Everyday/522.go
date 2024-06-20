package Everyday

import "sort"

// 判断 s 是否为 t 的子序列
func isSubseq(s, t string) bool {
    i := 0
    for _, c := range t {
        if s[i] == byte(c) {
            i++
            if i == len(s) { // 所有字符匹配完毕
                return true // s 是 t 的子序列
            }
        }
    }
    return false
}

func findLUSlength2(strs []string) int {
    ans := -1
    sort.Slice(strs, func(i, j int) bool {
        return len(strs[i]) > len(strs[j])
    })
next:
    for i, s := range strs {
        for j, t := range strs {
            if j != i && isSubseq(s, t) {
                continue next
            }
        }
        return len(strs[i])
    }
    return ans
}
