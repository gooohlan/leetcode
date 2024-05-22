package Everyday

import "slices"

func findMinimumTime(tasks [][]int) int {
    slices.SortFunc(tasks, func(a, b []int) int {
        return a[1] - b[1]
    })
    run := make([]bool, tasks[len(tasks)-1][1]+1)

    for _, task := range tasks {
        start, end, d := task[0], task[1], task[2]
        for _, b := range run[start : end+1] {
            if b {
                d--
            }
        }

        for i := end; d > 0; i-- {

        }
    }

    return 0
}
