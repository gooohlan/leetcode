package main

// [1,2,3]
// [4,5,6]
//
// 1,2,3,4,5,6
//
// [1,2]
// [3,4]
// [5,6]

func matrixReshape(nums [][]int, r int, c int) [][]int {
	x, y := len(nums), len(nums[0])
	if r*c != x*y {
		return nums
	}
	i, j := 0, 0
	newNums := make([][]int, r)
	for _, row := range nums {
		for _, v := range row {
			newNums[i] = append(newNums[i], v)
			j++
			if j == c {
				j = 0
				i++
			}
		}
	}
	return newNums
}

func matrixReshape2(nums [][]int, r int, c int) [][]int {
	x, y := len(nums), len(nums[0])
	if r*c != x*y {
		return nums
	}
	newNums := make([][]int, r)
	for i := 0; i < r; i++ {
		newNums[i] = make([]int, c)
	}
	for i := 0; i < r*c; i++ {
		// 两个二维数组都是基于同一个一维数组变化而来,他们在一维数组中的值是相等的
		newNums[i/c][i%c] = nums[i/y][i%y]
	}
	return newNums
}

// 原矩阵元素个数len(nums)*len(nums[0]) col:=lecoln(nums[0])
// i 从0开始遍历nums
// 原数组nums 行数 i/len(col) 列数 i%len(col))
// 目标数组 res 行数 i/c 列数 i%c
