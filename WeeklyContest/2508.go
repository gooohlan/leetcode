package WeeklyContest

type pair struct {
    x, y int
}

func isPossible(n int, edges [][]int) bool {
    has := make(map[pair]bool)
    deg := make([]int, n+1)
    
    for _, edge := range edges {
        x, y := edge[0], edge[1]
        has[pair{x, y}] = true
        has[pair{y, x}] = true
        deg[x]++
        deg[y]++
    }
    odd := []int{} // 记录出度为奇数的节点
    for i, d := range deg {
        if d%2 != 0 {
            odd = append(odd, i)
        }
    }
    
    // 奇数点数为0,已经符合要求
    if len(odd) == 0 {
        return true
    }
    // 奇数点数为2
    if len(odd) == 2 {
        x, y := odd[0], odd[1]
        if !has[pair{x, y}] { // 两个奇数点未相连,直接相连即可
            return true
        }
        // 找到点i, i满足不与x,y有边,连接i->x,i->y即可
        for i := 1; i < n; i++ {
            if i != x && i != y && !has[pair{i, x}] && !has[pair{i, y}] {
                return true
            }
        }
        return false
    }
    
    // 4个奇数点,分别记作a,b,c,d, 存在两两不相连,连接即可
    if len(odd) == 4 {
        a, b, c, d := odd[0], odd[1], odd[2], odd[3]
        return !has[pair{a, b}] && !has[pair{c, d}] ||
            !has[pair{a, c}] && !has[pair{b, d}] ||
            !has[pair{a, d}] && !has[pair{c, b}]
    }
    
    return false
}
