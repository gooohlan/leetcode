package Array

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func TestRemoveElement(t *testing.T) {
	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	val := 2
	fmt.Println(removeElement(arr, val))
}
