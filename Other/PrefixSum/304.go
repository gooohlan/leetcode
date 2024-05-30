package PrefixSum

type NumMatrix [][]int

func Constructor(matrix [][]int) NumMatrix {
    m, n := len(matrix), len(matrix[0])
    sum := make([][]int, m+1)
    sum[0] = make([]int, n+1)
    for i, row := range matrix {
        sum[i+1] = make([]int, n+1)
        for j, num := range row {
            sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + num
        }
    }
    return sum
}

func (n NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    return n[row2+1][col2+1] - n[row2+1][col1] - n[row1][col2+1] + n[row1][col1]
}
