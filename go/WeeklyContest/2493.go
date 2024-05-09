package WeeklyContest

func magnificentSets(n int, edges [][]int) (res int) {
    var (
        graph = make([][]int, n)
        bfs   func(int, int) int
        time  = make([]int, n) // 充当 visited 数组的作用，由于要重复 BFS，用 int 表示
        clock int              // 记录当前bfs遍历层数
    )
    for _, edge := range edges {
        graph[edge[0]-1] = append(graph[edge[0]-1], edge[1]-1)
        graph[edge[1]-1] = append(graph[edge[1]-1], edge[0]-1)
    }
    
    bfs = func(start, base int) int {
        var maxM int
        clock++
        time[start] = clock
        type queue struct{ start, base int }
        q := []queue{{start, base}}
        for len(q) > 0 {
            cur := q[0]
            q = q[1:]
            maxM = max(maxM, cur.base)
            for _, t := range graph[cur.start] {
                if time[t] != clock { // 不在同一层BFS中访问过
                    time[t] = clock
                    q = append(q, queue{t, cur.base + 1})
                }
            }
        }
        return maxM
    }
    
    colors := make([]int8, n)
    var nodes []int // 存储本轮已经染色的字符
    var isBipartite func(int, int8) bool
    isBipartite = func(i int, color int8) bool {
        nodes = append(nodes, i)
        colors[i] = color
        for _, t := range graph[i] {
            if colors[t] == color || colors[t] == 0 && !isBipartite(t, -color) { // 如果颜色一样活着给子节点燃相反颜色失败
                return false
            }
        }
        return true
    }
    
    for i, color := range colors {
        if color == 0 { // 此节点未染色
            nodes = nil
            // 进一步地，|y-x|=1 这个要求，也可以看成是从一点出发，通过两条不同路径到达另一个点时，这两条路径长度的奇偶性是相同的。
            // 换句话说，图中所有的环都必须有偶数个顶点，这等价于图是二分图
            if !isBipartite(i, 1) { // 不是二分图(有奇环)
                return -1
            }
            // 否则一定可以分组
            base := res + 1 // 接着上个连通分量的最大编号
            for _, node := range nodes {
                res = max(res, bfs(node, base))
            }
        }
    }
    return res
}

func abs(a, b int) int {
    if a > b {
        return a - b
    }
    return b - a
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}
