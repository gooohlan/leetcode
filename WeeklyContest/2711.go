package WeeklyContest

func differenceOfDistinctValues(grid [][]int) [][]int {
    m, n := len(grid), len(grid[0])
    ans := make([][]int, m)
    
    for i := range ans {
        ans[i] = make([]int, n)
        for j := range ans[i] {
            // 左上
            set := map[int]struct{}{}
            for x, y := i-1, j-1; x >= 0 && y >= 0; {
                set[grid[x][y]] = struct{}{}
                x--
                y--
            }
            sz := len(set)
            
            // 右下
            set = map[int]struct{}{}
            for x, y := i+1, j+1; x < m && y < n; {
                set[grid[x][y]] = struct{}{}
                x++
                y++
            }
            ans[i][j] = abs(sz, len(set))
        }
    }
    return ans
}

func differenceOfDistinctValues2(grid [][]int) [][]int {
    m, n := len(grid), len(grid[0])
    ans := make([][]int, m)
    for i := range ans {
        ans[i] = make([]int, n)
    }
    for s := 1; s < m+n; s++ {
        minJ := max(0, n-s)
        maxJ := min(n-1, n-s+m-1)
        
        // topLeft
        set := map[int]struct{}{}
        for j := minJ; j < maxJ; j++ {
            i := s + j - n
            set[grid[i][j]] = struct{}{}
            ans[i+1][j+1] = len(set) // 左上个数
        }
        
        // bottomRight
        set = map[int]struct{}{}
        for j := maxJ; j > minJ; j-- {
            i := s + j - n
            set[grid[i][j]] = struct{}{}
            ans[i-1][j-1] = abs(ans[i-1][j-1], len(set)) // ans[i-1][j-1]为左上个数, len(set)为右下个数,这里直接求解了
        }
    }
    return ans
}
