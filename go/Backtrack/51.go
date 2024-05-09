package Backtrack

var solutions [][]string

// 同一行同一列以及斜角位置不能出现皇后
// 因为是for循环,所以不会出现同一行,同一列我们加一个columns[i]的map即可
// 因为我们是从上往下放,所斜角只需要检测左上和右上
// 左上斜角满足行下标与列下标之差相等,比如 (0,0)和(2,2), (0,1)和(1,2)
// 右上斜角满足行下标与列下标之和相等,比如 (3,0)和(2,1), (3,1)和(2,2)
func solveNQueens(n int) [][]string {
    solutions = make([][]string, 0)
    queens := make([]int, n)
    for i := 0; i < n; i++ {
        queens[i] = -1
    }
    columns, diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}, map[int]bool{}
    backtrack51(queens, n, 0, columns, diagonals1, diagonals2)
    return solutions
}

func backtrack51(queens []int, n, row int, columns, diagonals1, diagonals2 map[int]bool) {
    // 触发结束条件
    if row == n {
        // 构造结果集
        board := generateBoard(queens, n)
        solutions = append(solutions, board)
        return
    }
    
    for i := 0; i < n; i++ {
        // 排除不合法选择
        if columns[i] {
            continue
        }
        if diagonals1[row-i] {
            continue
        }
        if diagonals2[row+i] {
            continue
        }
        // 做选择
        queens[row] = i
        columns[i] = true
        diagonals1[row-i] = true
        diagonals2[row+i] = true
        // 进入下一行决策
        backtrack51(queens, n, row+1, columns, diagonals1, diagonals2)
        // 撤销选择
        queens[row] = -1
        columns[i] = false
        diagonals1[row-i] = false
        diagonals2[row+i] = false
    }
    
}

func generateBoard(queens []int, n int) []string {
    board := []string{}
    for i := 0; i < n; i++ {
        row := make([]byte, n)
        for j := 0; j < n; j++ {
            row[j] = '.'
        }
        row[queens[i]] = 'Q'
        board = append(board, string(row))
    }
    return board
}
