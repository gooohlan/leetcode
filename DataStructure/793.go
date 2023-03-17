package DataStructure

import "sort"

func preimageSizeFZF(k int) int {
    return help2(k+1) - help2(k)
}

func help2(k int) int {
    return sort.Search(5*k, func(i int) bool {
        return trailingZeroes(i) >= k
    })
}

func help(target int) int {
    lo, hi := 0, 5*target
    for lo <= hi {
        mid := (hi-lo)/2 + lo
        if trailingZeroes(mid) < target {
            lo = mid + 1
        } else {
            hi = mid - 1
        }
    }
    return lo
}
