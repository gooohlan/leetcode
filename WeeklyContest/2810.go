package WeeklyContest

func finalString(s string) string {
    q := make([]rune, 0)
    tail := true
    for _, v := range s {
        if v == 'i' {
            tail = !tail
        } else if tail {
            q = append(q, v)
        } else {
            q = append([]rune{v}, q...)
        }
    }
    if !tail {
        for i, j := 0, len(q); i < j/2; i++ {
            q[i], q[j-1-i] = q[j-1-i], q[i]
        }
    }
    return string(q)
}

// 奇数次或者偶数次的翻转的字符是相连的,将他们分别放到两个数组中
// 翻转 (count(i)+1)/2 部分并放在前面,连接剩下的部分就OK了
func finalString2(s string) string {
    qs := [2][]rune{}
    tail := 1
    for _, v := range s {
        if v == 'i' {
            tail ^= 1
        } else {
            qs[tail] = append(qs[tail], v)
        }
    }
    q := qs[tail^1]
    for i, n := 0, len(q); i < n/2; i++ {
        q[i], q[n-1-i] = q[n-1-i], q[i]
    }
    return string(append(q, qs[tail]...))
}
// abefjk
// cdgh
0-翻转1
1 翻转0
2-翻转1
3-翻转0