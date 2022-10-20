package Array

import (
	"fmt"
	"testing"
)

func Test303(t *testing.T) {
	nums := Constructor([]int{-2, 0, 3, -5, 2, -1})

	fmt.Println(nums.SumRange(0, 2))
	fmt.Println(nums.SumRange(2, 5))
	fmt.Println(nums.SumRange(0, 5))
}

func Test304(t *testing.T) {
	nums := [][]int{
		[]int{3, 0, 1, 4, 2},
		[]int{5, 6, 3, 2, 1},
		[]int{1, 2, 0, 1, 5},
		[]int{4, 1, 0, 1, 7},
		[]int{1, 0, 3, 0, 5}}
	matrix := ConstructorMatrix(nums)
	fmt.Println(matrix.SumRegion(2, 1, 4, 3))
	fmt.Println(matrix.SumRegion(1, 1, 2, 2))
	fmt.Println(matrix.SumRegion(1, 2, 2, 4))
}