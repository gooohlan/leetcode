package Array

// 还是按照54题思路来,不过这次我们是放而不是取
func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	num := 1
	left, right, top, bottom := 0, n-1, 0, n-1
	for num <= n*n {

		// 左上向右
		if top <= bottom {
			for column := left; column <= right; column++ {
				matrix[top][column] = num
				num++
			}
			top++
		}

		// 右上到下
		if left <= right {
			for row := top; row <= bottom; row++ {
				matrix[row][right] = num
				num++
			}
			right--
		}

		// 右下到左
		if top <= bottom {
			for column := right; column >= left; column-- {
				matrix[bottom][column] = num
				num++
			}
			bottom--
		}
		// 左下到上
		if left <= right {
			for row := bottom; row >= top; row-- {
				matrix[row][left] = num
				num++
			}
			left++
		}
	}
	return matrix
}
