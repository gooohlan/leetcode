package Backtracking

import (
	"fmt"
	"testing"
)

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}
	list := permute(nums)
	fmt.Println(list)
}
