package Backtrack

func solveSudoku(board [][]byte) {
    var (
        m, n      = 9, 9
        backtrack func(row, col int) bool
        isValid   func(row, col int, n byte) bool
    )
    
    isValid = func(row, col int, n byte) bool {
        for i := 0; i < 9; i++ {
            // 判断行是否存在重复
            if board[row][i] == n {
                return false
            }
            // 判断列是否存在重复
            if board[i][col] == n {
                return false
            }
            // 判断 3 x 3 方框是否存在重复
            // row/3 和clo/3为3*3方格左上角
            if board[(row/3)*3+i/3][(col/3)*3+i%3] == n {
                return false
            }
        }
        return true
    }
    
    backtrack = func(row, col int) bool {
        if col == n {
            return backtrack(row+1, 0)
        }
        if row == m {
            return true
        }
        
        if board[row][col] != '.' {
            return backtrack(row, col+1)
        }
        
        for i := byte('1'); i <= '9'; i++ {
            if !isValid(row, col, i) {
                continue
            }
            board[row][col] = i
            if backtrack(row, col+1) {
                return true
            }
            board[row][col] = '.' // 还原
        }
        return false
    }
    
    backtrack(0, 0)
}
