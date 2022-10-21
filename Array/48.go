package Array

// 先将数组沿着 左上到右下的对角线旋转
// 然后从左往右翻转下,就得到了我们想要的
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j] // 对角线旋转
		}
		for j, k := 0, n-1; j < k; {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j] // 左右翻转
			j++
			k--
		}
	}

	return
}
