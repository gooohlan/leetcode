package Backtrack

func totalNQueens(n int) int {
    columns, diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}, map[int]bool{}
    var count int
    var backtrack func(int)
    backtrack = func(row int) {
        if row == n {
            count++
        }
        
        for i := 0; i < n; i++ {
            if columns[i] || diagonals1[row-i] || diagonals2[row+i] {
                continue
            }
            
            columns[i] = true
            diagonals1[row-i] = true
            diagonals2[row+i] = true
            backtrack(row + 1)
            columns[i] = false
            diagonals1[row-i] = false
            diagonals2[row+i] = false
        }
    }
    
    backtrack(0)
    return count
}
