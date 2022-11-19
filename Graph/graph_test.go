package Graph

import (
	"fmt"
	"testing"
)

func Test797(t *testing.T) {
	graph := [][]int{
		[]int{1, 2},
		[]int{3},
		[]int{3},
		[]int{},
	}
	res := allPathsSourceTarget(graph)
	fmt.Println(res)
}
