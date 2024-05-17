package Everyday

import "slices"

func numberOfWeeks(milestones []int) int64 {
    rest := 0
    for _, m := range milestones {
        rest += m
    }

    longest := slices.Max(milestones)
    rest -= longest
    if longest > rest+1 {
        return int64(rest*2 + 1)
    }
    return int64(longest + rest)
}
