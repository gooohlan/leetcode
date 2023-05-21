package WeeklyContest

// 先寻路,如果路上存在-1,则修改
func modifiedGraphEdges(n int, edges [][]int, source int, destination int, target int) [][]int {
    type Edge struct {
        From   int
        To     int
        Weight int
    }
    
    // var res [][]Edge
    
    graph := make([][]int, n)
    for i := range graph {
        graph[i] = make([]int, n)
    }
    for _, edge := range edges {
        to := edge[0]
        from := edge[1]
        weight := edge[2]
        graph[to][from] = weight
        graph[from][to] = weight
    }
    visited := make([]bool, n)
    var dfs func(source, destination int, path []int)
    var routes [][]int
    path := []int{source}
    
    dfs = func(source, target int, path []int) {
        if target == destination {
            routes = append(routes, path)
            return
        }
        visited[target] = true
        path = append(path, target)
        for from := range graph[target] {
            if !visited[from] {
                dfs(target, from, path)
            }
        }
        visited[target] = false
        path = path[:len(path)-1]
    }
    
    for form := range graph[source] {
        if graph[0][form] != 0 {
            dfs(source, form, path)
        }
    }
    
    var res [][]int
    for _, route := range routes {
        wight := 0
        indexs := []int{}
        for i := 1; i < len(route); i++ {
            w := graph[route[i-1]][route[i]]
            wight += w
            if w < 0 {
                indexs = append(indexs, i)
            }
            res = append(res, []int{route[i-1], route[i], w})
        }
        if len(indexs) == 0 && wight == target {
            return res
        }
        
        if target-wight < len(indexs) || wight > target {
            continue
        }
        
        for _, index := range indexs {
            res[index][2] = 1
        }
        res[len(indexs)-1][2] += target - wight - len(indexs)
        return res
    }
    return [][]int{}
}
