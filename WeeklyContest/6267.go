package WeeklyContest

func isPossible(n int, edges [][]int) bool {
    var (
        graph    = make([][]int, n+1)
        indegree = make([]int, n+1) // 入度数
    )
    
    for _, item := range edges {
        graph[item[1]] = append(graph[item[1]], item[0])
        graph[item[0]] = append(graph[item[0]], item[1])
        indegree[item[0]]++
    }
    
    oddEdge := make(map[int]bool)
    for i := 1; i < n+1; i++ {
        if indegree[i]%2 != 0 {
            oddEdge[i] = true
        }
    }
    if len(oddEdge) > 4 {
        return false
    }
    res := 0
    for v, _ := range oddEdge {
        i := 1
        for _, cur := range graph[v] {
            for ; i < cur; i++ {
                if oddEdge[i] && i != v {
                    res++
                    delete(oddEdge, v)
                    delete(oddEdge, i)
                    break
                }
            }
            if !oddEdge[v] {
                break
            }
            i++
        }
    }
    if res > 2 || len(oddEdge) != 0 {
        return false
    }
    return true
}
