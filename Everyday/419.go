package Everyday

func countBattleships(board [][]byte) int {
    ans := 0
    for i, b := range board {
        for j, b2 := range b {
            if b2 == 'X' && (i == 0 || board[i-1][j] == '.') && (j == 0 || board[i][j-1] == '.') {
                ans++
            }
        }
    }
    return ans
}
