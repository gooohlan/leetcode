package WeeklyContest

func similarPairs(words []string) int {
    arr := make([]int, len(words))
    
    for i, word := range words {
        vis := make(map[int32]bool)
        for _, w := range word {
            if !vis[w] {
                vis[w] = true
                arr[i] += int(w)
            }
        }
    }
    res := 0
    for i := 0; i < len(words); i++ {
        for j := i + 1; j < len(words); j++ {
            if arr[i] == arr[j] {
                res++
            }
        }
    }
    return res
}
func similarPairs2(words []string) (ans int) {
    cnt := map[int]int{}
    for _, s := range words {
        mask := 0
        for _, c := range s {
            mask |= 1 << (c - 'a')
        }
        ans += cnt[mask]
        cnt[mask]++
    }
    return
}
