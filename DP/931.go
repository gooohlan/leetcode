package DP

func minFallingPathSum(matrix [][]int) int {
    var dp func(int, int) int
    m, n := len(matrix), len(matrix[0])
    memo := make([][]int, m)
    for i := range memo {
        memo[i] = make([]int, n)
        for j := 0; j < n; j++ {
            memo[i][j] = 10001 // 数组最大为100*100, 数组值最大为100
        }
    }
    
    dp = func(i int, j int) int {
        if i < 0 || j < 0 || i >= m || j >= n {
            return 10001
        }
        
        if i == m-1 { // 到达最后一行
            return matrix[i][j]
        }
        
        if memo[i][j] == 10001 {
            memo[i][j] = matrix[i][j] + min(
                dp(i+1, j),
                dp(i+1, j-1),
                dp(i+1, j+1),
            )
        }
        return memo[i][j]
    }
    
    res := 10001
    // 起点可能是第一排的任何一列
    for j := 0; j < n; j++ {
        res = min(res, dp(0, j))
    }
    return res
}

func minFallingPathSumDP(matrix [][]int) int {
    n := len(matrix)
    for i := 1; i < n; i++ {
        for j := 0; j < n; j++ {
            if j == 0 {
                matrix[i][j] += min(matrix[i-1][j], matrix[i-1][j+1])
            } else if j == n-1 {
                matrix[i][j] += min(matrix[i-1][j], matrix[i-1][j-1])
            } else {
                matrix[i][j] += min(matrix[i-1][j], matrix[i-1][j-1], matrix[i-1][j+1])
            }
        }
    }
    res := matrix[n-1][0]
    for i := 1; i < n; i++ {
        if res > matrix[n-1][i] {
            res = matrix[n-1][i]
        }
    }
    
    return res
}
