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

func Test207(t *testing.T) {
	fmt.Println(canFinishBFS(2, [][]int{[]int{1, 0}}))
}

func Test210(t *testing.T) {
	res := findOrderBFS(4, [][]int{
		[]int{1, 0},
		[]int{2, 0},
		[]int{3, 1},
		[]int{3, 2},
	})
	fmt.Println(res)
}
