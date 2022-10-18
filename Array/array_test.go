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
