package WeeklyContest

func maxOutput(n int, edges [][]int, price []int) int64 {
    var (
        graph = make([][]int, n)
        dfs   func(int, int) (int, int) // 返回带叶子的最大路径和，不带叶子的最大路径和
        res   int
    )
    
    for _, item := range edges {
        graph[item[1]] = append(graph[item[1]], item[0])
        graph[item[0]] = append(graph[item[0]], item[1])
    }
    graph[0] = append(graph[0], -1) // 防止根节点被认作叶子
    dfs = func(x, father int) (maxS1, maxS2 int) {
        price := price[x]
        if len(graph[x]) == 1 { // 叶子节点
            return price, 0
        }
        
        for _, y := range graph[x] {
            if y != father {
                s1, s2 := dfs(y, x)
                if maxS2 == 0 { // 特判：第一颗子树，这个时候不能选当前节点
                    res = max(res, s1)
                } else { // 前面最大不带叶子的路径 + 当前节点 + 当前带叶子的路径
                    res = max(res, maxS2+price+s1)
                }
                // 前面最大带叶子的路径 + 当前节点 + 当前不带叶子的路径
                res = max(res, maxS1+price+s2)
                maxS1 = max(maxS1, s1)
                maxS2 = max(maxS2, s2)
            }
            
        }
        return maxS1 + price, maxS2 + price
    }
    
    dfs(0, -1)
    return int64(res)
}
