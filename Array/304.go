package Array

type NumMatrix struct {
	sumMatrix [][]int
}

func ConstructorMatrix(matrix [][]int) NumMatrix {
	m := len(matrix)
	n := len(matrix[0])
	sumMatrix := make([][]int, m+1)
	sumMatrix[0] = make([]int, n+1)
	for i, row := range matrix {
		sumMatrix[i+1] = make([]int, n+1)
		for j, v := range row {
			sumMatrix[i+1][j+1] = sumMatrix[i+1][j] + sumMatrix[i][j+1] + v - sumMatrix[i][j]
		}
	}
	return NumMatrix{
		sumMatrix: sumMatrix,
	}
}

func (nums *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return nums.sumMatrix[row2+1][col2+1] - nums.sumMatrix[row1][col2+1] - nums.sumMatrix[row2+1][col1] + nums.sumMatrix[row1][col1]

}
