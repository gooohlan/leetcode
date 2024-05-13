package Everyday

type pair struct{ x, y int }

var dis = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func orangesRotting(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    fresh := 0
    queue := []pair{}

    for i, row := range grid {
        for j, orange := range row {
            if orange == 1 {
                fresh++
            } else if orange == 2 {
                queue = append(queue, pair{i, j})
            }
        }
    }

    res := -1
    for len(queue) > 0 {
        res++
        tmp := queue
        queue = []pair{}

        for _, p := range tmp {
            for _, d := range dis {
                i, j := p.x+d.x, p.y+d.y
                if i >= 0 && i < m && j >= 0 && j < n && grid[i][j] == 1 {
                    grid[i][j] = 2
                    fresh--
                    queue = append(queue, pair{i, j})
                }
            }
        }
    }

    if fresh > 0 {
        return -1
    }

    return max(res, 0)
}
