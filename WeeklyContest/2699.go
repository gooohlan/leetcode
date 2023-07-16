package WeeklyContest

import "math"

func modifiedGraphEdges(n int, edges [][]int, source int, destination int, target int) [][]int {
    type edge struct {
        to, eid int // eid方便后续修改edges
    }
    graph := make([][]edge, n)
    for i, e := range edges {
        form, to := e[0], e[1]
        graph[form] = append(graph[form], edge{to, i})
        graph[to] = append(graph[to], edge{form, i})
    }
    
    var delta int
    dis := make([][2]int, n)
    for i := range dis {
        dis[i] = [2]int{math.MaxInt, math.MaxInt}
    }
    dis[source] = [2]int{}
    dijkstra := func(k int) {
        vis := make([]bool, n)
        for {
            // 找到最短路,并且没有vis的点,(第一轮循环找的是起点 source)
            x := -1
            for y, b := range vis {
                if !b && (x < 0 || dis[y][k] < dis[x][k]) {
                    x = y
                }
            }
            
            if x == destination { // 起点source到终点 destination的最短距离已确认
                return
            }
            vis[x] = true
            for _, e := range graph[x] {
                y, wt := e.to, edges[e.eid][2]
                if wt == -1 {
                    wt = 1 // -1改为1
                }
                if k == 1 && edges[e.eid][2] == -1 { // 第二次 Dijkstra，改成 w
                    w := delta + dis[y][0] - dis[x][1]
                    if w > wt {
                        wt = w
                        edges[e.eid][2] = w
                    }
                }
                // 更新最短路
                dis[y][k] = min(dis[y][k], dis[x][k]+wt)
            }
        }
    }
    
    // 第一次dijkstra
    dijkstra(0)
    delta = target - dis[destination][0]
    if delta < 0 { // 起点到终点长度好过target,无解,  -1 全改为 1 时，最短路比 target 还大
        return nil
    }
    
    // 第二次dijkstra
    dijkstra(1)
    if dis[destination][1] < target { // 如果第二次dijkstra,最短路无法再变大，无法达到 target
        return nil
    }
    
    for _, e := range edges {
        if e[2] == -1 { // 剩余没修改的边全部改成 1
            e[2] = 1
        }
    }
    return edges
}
