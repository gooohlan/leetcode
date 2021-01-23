package main

import "log"

func main() {
	arr := [][]int{
		[]int{1, 1, 1},
		[]int{1, 0, 1},
		[]int{1, 1, 1},
	}
	newArr := imageSmoother(arr)
	log.Println(newArr)
}

func imageSmoother(M [][]int) [][]int {
	var x, y = len(M), len(M[0])
	newM := make([][]int, x)
	for i := 0; i < x; i++ {
		newM[i] = make([]int, y)
	}
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			var count int
			// 寻找邻居格子
			for ii := i-1; ii <= i+1; ii++ {
				for jj := j-1; jj <= j+1; jj++ {
					if ii >=0 && ii < x && jj >=0 && jj < y { // 去除超出原数组范围的数据
						newM[i][j] += M[ii][jj]
						count ++
					}
				}
			}
			newM[i][j] /= count
		}
	}
	return newM
}
