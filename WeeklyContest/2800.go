package WeeklyContest

import "strings"

func merge(s, t string) string {
    // 先特判完全包含的情况
    if strings.Contains(s, t) {
        return s
    }
    if strings.Contains(t, s) {
        return t
    }
    for i := min(len(s), len(t)); ; i-- {
        // 枚举：s 的后 i 个字母和 t 的前 i 个字母是一样的
        x := s[len(s)-i:]
        y := t[:i]
        if x == y {
            return s + t[i:]
        }
    }
}

func minimumString(a string, b string, c string) (res string) {
    strList := []string{a, b, c}
    
    arr := [][]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 1, 0}, {2, 0, 1}}
    for _, a := range arr {
        s := merge(merge(strList[a[0]], strList[a[1]]), strList[a[2]])
        if res == "" || len(s) < len(res) || len(s) == len(res) && s < res {
            res = s
        }
    }
    return
}
