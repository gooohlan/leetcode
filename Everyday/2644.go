package Everyday

import "sort"

func maxDivScore(nums []int, divisors []int) int {
    res := 0
    maxCnt := -1
    for _, d := range divisors {
        cnt := 0
        for _, n := range nums {
            if n%d == 0 {
                cnt++
            }
        }
        if cnt > maxCnt || cnt == maxCnt && d < res {
            maxCnt = cnt
            res = d
        }
    }
    return res
}

func maxDivScore2(nums []int, divisors []int) int {
    res := 0
    maxCnt := -1
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] > nums[j]
    })
    for _, d := range divisors {
        cnt := 0
        for _, n := range nums {
            if n < d {
                break
            }
            if n%d == 0 {
                cnt++
            }
        }
        if cnt > maxCnt || cnt == maxCnt && d < res {
            maxCnt = cnt
            res = d
        }
    }
    return res
}
