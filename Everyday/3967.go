package Everyday

type edge struct {
    to, weight int
}

func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
    n := len(edges) + 1
    g := make([][]edge, n)
    for _, e := range edges {
        u, v, w := e[0], e[1], e[2]
        g[u] = append(g[u], edge{v, w})
        g[v] = append(g[v], edge{u, w})
    }
    
    var dfs func(int, int, int) int
    dfs = func(to, from, sum int) int {
        cnt := 0
        if sum%signalSpeed == 0 {
            cnt = 1
        }
        for _, e := range g[to] {
            if e.to != from {
                cnt += dfs(e.to, to, sum+e.weight)
            }
        }
        return cnt
    }
    
    res := make([]int, n)
    for i, gi := range g {
        if len(gi) == 1 {
            continue
        }
        s := 0
        for _, e := range gi {
            cnt := dfs(e.to, i, e.weight)
            res[i] += cnt * s
            s += cnt
        }
    }
    return res
}
