package WeeklyContest

import "sort"

func maximumCount(nums []int) int {
    a, b := 0, 0
    for _, num := range nums {
        if num < 0 {
            a++
        }
        if num > 0 {
            b++
        }
    }
    return max(a, b)
}

func maximumCount2(nums []int) int {
    return max(sort.SearchInts(nums, 0), len(nums)-sort.SearchInts(nums, 1))
}
