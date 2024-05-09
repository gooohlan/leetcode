package WeeklyContest

// 赋值以后面的为主,所以倒序赋值
// 赋值时如果当前行没被赋值过,那么就进行赋值,如果赋值过,就不操作
// 同时,如果该行的某几列被复制过,那么就不进行赋值,所以需要记录赋值过的列
func matrixSumQueries(n int, queries [][]int) int64 {
    vis := [2]map[int]bool{{}, {}} // 0: 行, 1: 列
    var ans int64
    for i := len(queries) - 1; i >= 0; i-- {
        t, index, val := queries[i][0], queries[i][1], queries[i][2]
        if !vis[t][index] { // 没有被赋值过
            var remain int
            if t == 0 {
                remain = n - len(vis[1])
            } else {
                remain = n - len(vis[0])
            }
            // 这一行还剩下 remain 个可以填入的格子
            ans += int64(remain * val)
            vis[t][index] = true
        }
    }
    return ans
}
