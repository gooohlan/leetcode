package Everyday

func minimumSteps(s string) int64 {
    ans := 0
    cnt := 0
    for _, c := range s {
        if c == '1' {
            cnt++
        } else {
            ans += cnt
        }
    }
    return int64(ans)
}
