package Everyday

func minAnagramLength(s string) int {
    n := len(s)
    for k := 1; k <= n/2; k++ {
        if n%k > 0 {
            continue
        }
        cnt0 := [26]int{}
        for _, b := range s[:k] {
            cnt0[b-'a']++
        }
        find := true
        for i := k * 2; i <= len(s); i += k {
            cnt := [26]int{}
            for _, b := range s[i-k : i] {
                cnt[b-'a']++
            }
            if cnt != cnt0 {
                find = false
                break
            }
        }
        if find {
            return k
        }
    }
    return n
}
