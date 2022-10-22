package Array

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	rows, columns := len(matrix), len(matrix[0])
	left, right, top, bottom := 0, columns-1, 0, rows-1

	res := make([]int, rows*columns)
	var index int
	for left <= right && top <= bottom {
		// 左上向右
		if top <= bottom {
			for column := left; column <= right; column++ {
				res[index] = matrix[top][column]
				index++
			}
			top++
		}
		// 右上到下
		if left <= right {
			for row := top; row <= bottom; row++ {
				res[index] = matrix[row][right]
				index++
			}
			right--
		}
		// 右下到左
		if top <= bottom {
			for column := right; column >= left; column-- {
				res[index] = matrix[bottom][column]
				index++
			}
			bottom--
		}
		// 左下到上
		if left <= right {
			for row := bottom; row >= top; row-- {
				res[index] = matrix[row][left]
				index++
			}
			left++
		}
	}
	return res
}
