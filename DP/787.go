package DP

import (
    "container/heap"
    "math"
)

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

// DP解法
func findCheapestPriceDP(n int, flights [][]int, src int, dst int, k int) int {
    dp := make([][]int, k+2) // dp[t][j]表示遍历t条边到达目的地j的最短路径
    for i := range dp {
        dp[i] = make([]int, n)
        for j := 0; j < n; j++ {
            dp[i][j] = math.MaxInt / 2
        }
    }
    dp[0][src] = 0
    
    res := math.MaxInt / 2
    for t := 1; t < k+2; t++ {
        for _, flight := range flights {
            i, j, cost := flight[0], flight[1], flight[2]
            dp[t][j] = min(dp[t][j], dp[t-1][i]+cost)
            if j == dst { // 到达终点
                res = min(res, dp[t][j])
            }
        }
    }
    if res == math.MaxInt/2 {
        return -1
    }
    return res
}

// DP解法-压缩
func findCheapestPriceDP2(n int, flights [][]int, src int, dst int, k int) int {
    dp := make([]int, n) //
    for i := range dp {
        dp[i] = math.MaxInt / 2
    }
    dp[src] = 0
    
    res := math.MaxInt / 2
    for t := 1; t < k+2; t++ {
        pre := make([]int, n)
        for i := range dp {
            pre[i] = math.MaxInt / 2
        }
        for _, flight := range flights {
            i, j, cost := flight[0], flight[1], flight[2]
            // dp[t][j] = min(dp[t][j], dp[t-1][i]+cost)
            pre[j] = min(pre[j], dp[i]+cost)
        }
        dp = pre
        res = min(res, dp[dst])
    }
    if res == math.MaxInt/2 {
        return -1
    }
    return res
}

// dijkstra 解法
// https://leetcode.cn/problems/cheapest-flights-within-k-stops/solution/c-kzhan-zhong-zhuan-nei-zui-bian-yi-de-h-ou4d/
func findCheapestPriceDijkstra(n int, flights [][]int, src int, dst int, k int) int {
    type edge struct { // 记录指向他的节点
        to, price int
    }
    graph := make([][]edge, n)
    
    for _, flight := range flights {
        graph[flight[0]] = append(graph[flight[0]], edge{flight[1], flight[2]})
    }
    k++
    
    costs := make([]int, n)    // 记录到指定节点最低消费
    nodeNums := make([]int, n) // 记录到当前节点最小步数
    for i := range costs {
        costs[i] = math.MaxInt
        nodeNums[i] = math.MaxInt
    }
    costs[src] = 0    // 起点消费为0
    nodeNums[src] = 0 // 起点步数为0
    
    pq := &hp{{src, 0, 0}}
    for pq.Len() > 0 {
        cur := heap.Pop(pq).(state)
        if cur.id == dst { // 找到最短路径,返回总消费
            return cur.cost
        }
        
        if cur.nodeNum == k { // 中转次数耗尽
            continue
        }
        
        // 将cur能访问的节点加入队列
        for _, next := range graph[cur.id] {
            nextCost := next.price + cur.cost                                 // 到达下个节点的消费加上当前节点总消费
            nextNodeNum := cur.nodeNum + 1                                    // 中转次数消耗+1
            if costs[next.to] < nextCost && nodeNums[next.to] < nextNodeNum { // 剪枝，如果中转次数更多，花费还更大，那必然不会是最短路径
                continue
            }
            costs[next.to] = nextCost       // 更新节点最小消耗
            nodeNums[next.to] = nextNodeNum // 更新节点最小消耗
            heap.Push(pq, state{next.to, nextCost, nextNodeNum})
        }
    }
    return -1
}

type state struct {
    id      int // 图节点id
    cost    int // 从src到当前节点的消费
    nodeNum int // 经过的节点数
}
type hp []state

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].cost < h[j].cost }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x interface{})   { *h = append(*h, x.(state)) }
func (h *hp) Pop() (v interface{}) { old := *h; *h, v = old[:len(old)-1], old[len(old)-1]; return }
