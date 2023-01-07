package DP

import "math"

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    type edge struct { // 记录指向他的节点
        from, price int
    }
    indegree := make([][]edge, n)
    for _, flight := range flights {
        indegree[flight[1]] = append(indegree[flight[1]], edge{flight[0], flight[2]})
    }
    
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, k+2)
    }
    
    // 定义：从 src 出发，j 步之内到达 i 的最短路径权重
    var dfs func(i, j int) int
    dfs = func(i, j int) int {
        // 从 src 到 src，一步都不用走
        if i == src {
            return 0
        }
        
        // 如果步数用尽，就无解了
        if j == 0 {
            return -1
        }
        
        if memo[i][j] == 0 {
            res := math.MaxInt
            for _, e := range indegree[i] {
                // 从 src 到达相邻的入度节点所需的最短路径权重
                subProblem := dfs(e.from, j-1)
                if subProblem != -1 { // 跳过无解
                    res = min(res, subProblem+e.price)
                }
            }
            if res == math.MaxInt {
                res = -1
            }
            memo[i][j] = res
        }
        
        return memo[i][j]
    }
    
    // 将中转站个数转化成边的条数
    k++
    
    return dfs(dst, k)
}

// dijkstra 解法
func findCheapestPrice2(n int, flights [][]int, src int, dst int, k int) int {
    type edge struct { // 记录指向他的节点
        to, price int
    }
    graph := make([][]edge, n)
    
    for _, flight := range flights {
        graph[flight[0]] = append(graph[flight[0]], edge{flight[1], flight[2]})
    }
    k++
    
}
