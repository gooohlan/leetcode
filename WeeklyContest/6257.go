package WeeklyContest

import (
	"math"
	"sort"
)

func deleteGreatestValue(grid [][]int) int {
	for _, ints := range grid {
		sort.Slice(ints, func(i, j int) bool {
			return ints[i] > ints[j]
		})
	}
	var res int
	for i := 0; i < len(grid[0]); i++ {
		max := math.MinInt
		for j := 0; j < len(grid); j++ {
			if grid[j][i] > max {
				max = grid[j][i]
			}
		}
		res += max
	}
	return res
}
