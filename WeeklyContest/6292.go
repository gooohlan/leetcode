package WeeklyContest

func rangeAddQueries(n int, queries [][]int) [][]int {
    arr := make([][]int, n)
    for i := range arr {
        arr[i] = make([]int, n)
    }
    
    for _, query := range queries {
        row1, col1, row2, col2 := query[0], query[1], query[2], query[3]
        for i := row1; i <= row2; i++ {
            arr[i][col1] += 1
            if col2 < n-1 {
                arr[i][col2+1] -= 1
            }
        }
    }
    for i := 0; i < n; i++ {
        for j := 1; j < n; j++ {
            arr[i][j] += arr[i][j-1]
        }
    }
    return arr
}

func rangeAddQueries2(n int, queries [][]int) [][]int {
    // m := n
    diff := make([][]int, n+1)
    for i := range diff {
        diff[i] = make([]int, n+1)
    }
    
    update := func(r1, c1, r2, c2, x int) {
        r2++
        c2++
        diff[r1][c1] += x
        diff[r1][c2] -= x
        diff[r2][c1] -= x
        diff[r2][c2] += x
    }
    
    for _, query := range queries {
        update(query[0], query[1], query[2], query[3], 1)
    }
    res := make([][]int, n+1)
    res[0] = make([]int, n+1)
    for i, row := range diff[:n] {
        res[i+1] = make([]int, n+1)
        for j, x := range row[:n] {
            res[i+1][j+1] = res[i+1][j] + res[i][j+1] - res[i][j] + x
        }
    }
    res = res[1:]
    for i, row := range res {
        res[i] = row[1:]
    }
    return res
}
