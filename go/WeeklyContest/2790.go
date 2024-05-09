package WeeklyContest

import "sort"

func maxIncreasingGroups(usageLimits []int) int {
    sort.Ints(usageLimits)
    curr, ans := 0, 0
    for _, n := range usageLimits {
        curr += n
        if curr >= ans+1 {
            ans += 1
            curr -= ans
        }
    }
    return ans
}
