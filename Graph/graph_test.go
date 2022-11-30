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

func Test130(t *testing.T) {
	board := [][]byte{
		[]byte{'X', 'O', 'X'},
		[]byte{'X', 'O', 'X'},
		[]byte{'X', 'O', 'X'},
	}
	solve(board)
	for _, bytes := range board {
		fmt.Println(string(bytes))
	}
}

func Test990(t *testing.T) {
	equations := []string{
		"c==c", "b==d", "id!=z",
	}
	fmt.Println(equationsPossible(equations))
}

func Test261(t *testing.T) {
	edges := [][]int{
		[]int{0, 1},
		[]int{1, 2},
		[]int{2, 3},
		[]int{1, 3},
		[]int{1, 4},
	}
	fmt.Println(validTreeDFS(5, edges))
	fmt.Println(validTreeBFS(5, edges))
	fmt.Println(validTreeUF(5, edges))
}

func Test1135(t *testing.T) {
	conections := [][]int{
		[]int{3, 4, 2},
		[]int{2, 3, 1},
		[]int{1, 4, 1},
	}
	fmt.Println(minimumCost(4, conections))
	fmt.Println(minimumCostPrim(4, conections))
}

func Test1584(t *testing.T) {
	points := [][]int{[]int{0, 0}, []int{2, 2}, []int{3, 10}, []int{5, 2}, []int{7, 0}}
	fmt.Println(minCostConnectPoints(points))
	fmt.Println(minCostConnectPointsPrim(points))
}

func Test743(t *testing.T) {
	times := [][]int{
		[]int{1, 2, 1},
		// []int{2, 3, 1},
		// []int{3, 4, 1},
	}

	fmt.Println(networkDelayTime(times, 2, 2))
}

func Test1631(t *testing.T) {
	heights := [][]int{
		[]int{1, 2, 2},
		[]int{3, 8, 2},
		[]int{5, 3, 5},
	}
	fmt.Println(minimumEffortPath(heights))
	fmt.Println(minimumEffortPathUF(heights))
}

func Test1514(t *testing.T) {
	edges := [][]int{
		[]int{0, 1},
	}

	succProb := []float64{0.5}

	fmt.Println(maxProbability(3, edges, succProb, 0, 2))
}
