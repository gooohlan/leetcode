package WeeklyContest

func countPalindromePaths(parent []int, s string) int64 {
    n := len(parent)
    type pair struct{ to, wt int }
    g := make([][]pair, n)
    for i := 1; i < n; i++ {
        p := parent[i]
        g[p] = append(g[p], pair{i, 1 << (s[i] - 'a')})
    }
    
    ans := 0
    cnt := map[int]int{0: 1} // 一条「空路径」
    var dfs func(int, int)
    dfs = func(v, xor int) {
        for _, e := range g[v] {
            x := xor ^ e.wt
            ans += cnt[x] // x ^ x = 0
            for i := 0; i < 26; i++ {
                ans += cnt[x^(1<<i)] // x ^ (x^(1<<i)) = 1<<i
            }
            cnt[x]++
            dfs(e.to, x)
        }
    }
    dfs(0, 0)
    return int64(ans)
}
