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

func Test785(t *testing.T) {
	graph := [][]int{
		[]int{1, 2, 3},
		[]int{0, 2},
		[]int{0, 1, 3},
		[]int{0, 2},
	}
	fmt.Println(isBipartite(graph))
	fmt.Println(isBipartiteBFS(graph))
}

func Test886(t *testing.T) {
	dislikes := [][]int{
		[]int{1, 2},
		[]int{1, 3},
		[]int{2, 3},
	}
	fmt.Println(possibleBipartition(3, dislikes))
}

func Test323(t *testing.T) {
	edges := [][]int{
		[]int{0, 1},
		[]int{1, 2},
		// []int{2, 3},
		[]int{3, 4},
	}
	fmt.Println(countComponents(5, edges))
	fmt.Println(countComponentsBfs(5, edges))
	fmt.Println(countComponentsBfsUF(5, edges))
}
