package Everyday

import (
    "sort"
)

func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
    n := len(difficulty)
    jobs := make([][]int, n)
    for i, d := range difficulty {
        jobs[i] = []int{d, profit[i]}
    }
    sort.Slice(jobs, func(i, j int) bool {
        return jobs[i][0] < jobs[j][0]
    })
    sort.Ints(worker)

    i, maxP, res := 0, 0, 0
    for _, w := range worker {
        for i < n && jobs[i][0] <= w { // 找到难度小于等于当前工人能力的工作
            maxP = max(maxP, jobs[i][1])
            i++ //
        }
        res += maxP
    }

    return res
}
